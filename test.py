import argparse
import importlib
from pathlib import Path
from types import ModuleType

from cimgen import cimgen


def build() -> None:
    outdir='output/python/CGMES_3.0.0'
    schemadir='cgmes_schema/CGMES_3.0.0'
    langdir='python'
    cgmes_version='cgmes_v3_0_0'

    lang_pack: ModuleType = importlib.import_module(f"cimgen.languages.{langdir}.lang_pack")
    schema_path = Path.cwd() / schemadir
    cimgen.cim_generate(schema_path, outdir, cgmes_version, lang_pack)


if __name__ == "__main__":
    build()
