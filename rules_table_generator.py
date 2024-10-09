#!/usr/bin/env python

from pathlib import Path
from urllib.parse import quote
import yaml
import sys

LANGUAGES = ['go', 'python', 'rs', 'javascript']
DIRS = ['go', 'optimizations', 'informational', 'experimental']
IMPACT_MAP = {
    'LOW': "🟩",
    'MEDIUM': "🟧",
    'HIGH': "🟥",
    None: "🌫️",
}
CONFIDENCE_MAP = {
    'LOW': "🌕",
    'MEDIUM': "🌗",
    'HIGH': "🌘",
    None: "",
}


def main():
    for dir in DIRS:
        rules_for_dir = []
        for rule_path in Path(dir).rglob('*.yaml'):
            try:
                rules_data = yaml.safe_load(rule_path.open())
            except yaml.YAMLError as err:
                print(f"Error reading {rule_path} - {err}", file=sys.stderr)
                continue

            if rules_data is None or 'rules' not in rules_data:
                print(f"Error for {rule_path} - missing rules", file=sys.stderr)
                continue

            rules_data = rules_data['rules']
            if len(rules_data) == 0:
                print(f"Error for {rule_path} - missing any rule", file=sys.stderr)
                continue

            for rule_data in rules_data:
                rules_for_dir.append((rule_path, rule_data))

        if len(rules_for_dir) > 0:
            print(f"### {dir}")
            print("")
            print("| ID | Impact | Confidence | Description |")
            print("| -- | :----: | :--------: | ----------- |")

            for rule_path, rule_data in sorted(rules_for_dir, key=lambda x: (x[0], x[1]['id'])):
                rule_meta = rule_data.get('metadata', {})
                print(
                    f"| [{rule_data['id']}]({rule_path}) | {IMPACT_MAP[rule_meta.get('impact')]} | {CONFIDENCE_MAP[rule_meta.get('confidence')]} | {rule_meta.get('description', '')} |"
                )

            print("\n")


if __name__ == "__main__":
    main()