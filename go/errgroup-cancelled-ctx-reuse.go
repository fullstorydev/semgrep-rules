package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()

	err := badUsage(ctx)
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = goodUsage(ctx)
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = beforeWaitIsOk(ctx)
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = reassignedContext(ctx)
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = multipleContextUsages(ctx)
	if err != nil {
		log.Fatal("Error:", err)
	}
}

// Simulated functions that take context
func fetchData(ctx context.Context) error {
	return nil
}

func processData(ctx context.Context, data string) error {
	return nil
}

func someFunc(ctx context.Context) error {
	return nil
}

func badUsage(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return fetchData(ctx)
	})

	err := g.Wait()
	if err != nil {
		return err
	}

	// ruleid: bad_errgroup
	err = processData(ctx, "some data")
	if err != nil {
		return err
	}

	return nil
}

func goodUsage(ctx context.Context) error {
	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return fetchData(gctx)
	})

	err := g.Wait()
	if err != nil {
		return err
	}

	// ok: bad_errgroup
	err = processData(ctx, "some data")
	if err != nil {
		return err
	}

	return nil
}

func beforeWaitIsOk(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return fetchData(ctx)
	})

	// ok: bad_errgroup
	err := processData(ctx, "some data")
	if err != nil {
		return err
	}

	err = g.Wait()
	if err != nil {
		return err
	}

	return nil
}

func reassignedContext(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return fetchData(ctx)
	})

	err := g.Wait()
	if err != nil {
		return err
	}

	// ok: bad_errgroup
	ctx = context.Background()
	err = processData(ctx, "some data")
	if err != nil {
		return err
	}

	return nil
}

func multipleContextUsages(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return fetchData(ctx)
	})

	g.Go(func() error {
		// ok: bad_errgroup
		return processData(ctx, "in goroutine")
	})

	err := g.Wait()
	if err != nil {
		return err
	}

	// ruleid: bad_errgroup
	err = processData(ctx, "after wait")
	if err != nil {
		return err
	}

	// ruleid: bad_errgroup
	err = someFunc(ctx)
	if err != nil {
		return err
	}

	return nil
}

func nestedErrgroups(ctx context.Context) error {
	g1, ctx := errgroup.WithContext(ctx)

	g1.Go(func() error {
		return fetchData(ctx)
	})

	err := g1.Wait()
	if err != nil {
		return err
	}

	g2, ctx := errgroup.WithContext(ctx)

	g2.Go(func() error {
		return fetchData(ctx)
	})

	err = g2.Wait()
	if err != nil {
		return err
	}

	// ruleid: bad_errgroup
	err = processData(ctx, "after second wait")
	if err != nil {
		return err
	}

	return nil
}

func differentVariableNames(ctx context.Context) error {
	eg, newCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return fetchData(newCtx)
	})

	err := eg.Wait()
	if err != nil {
		return err
	}

	// ruleid: bad_errgroup
	err = processData(newCtx, "using newCtx")
	if err != nil {
		return err
	}

	// ok: bad_errgroup
	err = processData(ctx, "using original ctx")
	if err != nil {
		return err
	}

	return nil
}

func contextInMultipleParams(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return fetchData(ctx)
	})

	err := g.Wait()
	if err != nil {
		return err
	}

	// ruleid: bad_errgroup
	fmt.Printf("Context: %v, Data: %s\n", ctx, "test")

	return nil
}

func earlyReturn(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return fetchData(ctx)
	})

	if err := g.Wait(); err != nil {
		// ruleid: bad_errgroup
		someFunc(ctx)
		return err
	}

	// ruleid: bad_errgroup
	err := processData(ctx, "after successful wait")
	return err
}

func withDeferredsomeFunc(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	defer func() {
		// ruleid: bad_errgroup
		someFunc(ctx)
	}()

	g.Go(func() error {
		return fetchData(ctx)
	})

	err := g.Wait()
	if err != nil {
		return err
	}

	// ruleid: bad_errgroup
	return processData(ctx, "before defer")
}
