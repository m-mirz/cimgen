import os
import chevron
import logging
import glob
from importlib.resources import files
from typing import Callable
from pathlib import Path
import shutil

logger = logging.getLogger(__name__)

# template files that are used to generate the code files
struct_template_file = {"filename": "golang_struct_template.mustache", "ext": ".go"}
constants_template_file = {"filename": "golang_enum_template.mustache", "ext": ".go"}
#profile_template_file = {"filename": "cimpy_profile_template.mustache", "ext": ".go"}
partials = {}

# Setup called only once: create clean output directory, base class, profile class, etc.
# cgmes_profile_details contains index, names and uris for each profile.
# We use that to create the header data for the profiles.
def setup(output_path: str, version: str, cgmes_profile_details: list[dict], namespaces: dict[str, str]) -> None:
    source_dir = Path(__file__).parent
    dest_dir = Path(output_path)
    for file in dest_dir.glob("**/*.java"):
        file.unlink()
    # Add all hardcoded utils and create parent dir
    for file in source_dir.glob("**/*.go"):
        dest_file = dest_dir / file.relative_to(source_dir)
        dest_file.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy(file, dest_file)
    
    #_create_constants(dest_dir, version, namespaces)
    #_create_cgmes_profile(output_path, cgmes_profile_details)


def get_base_class() -> str:
    return "Base"


def get_class_location(class_name: str, class_map: dict, version: str) -> str: # NOSONAR
    return ""


def run_template(output_path: str, class_details: dict) -> None:

    if class_details["is_a_primitive_class"]:
        return
    else:
        template = struct_template_file

    class_file = Path(output_path) / (class_details["class_name"] + template["ext"])
    _write_templated_file(class_file, class_details, template["filename"])


def _write_templated_file(class_file: Path, class_details: dict, template_filename: str) -> None:
    with class_file.open("w", encoding="utf-8") as file:        
        templates = files("cimgen.languages.golang.templates")
        with templates.joinpath(template_filename).open(encoding="utf-8") as f:
            args = {
                "data": class_details,
                "template": f,
                "partials_dict": partials,
            }
            output = chevron.render(**args)
        file.write(output)


def _create_constants(output_path: Path, version: str, namespaces: dict[str, str]) -> None:
    class_file = output_path / ("CimConstants" + constants_template_file["ext"])
    namespaces_list = [{"ns": ns, "uri": uri} for ns, uri in sorted(namespaces.items())]
    class_details = {"version": version, "namespaces": namespaces_list}
    _write_templated_file(class_file, class_details, constants_template_file["filename"])


class_blacklist = ["CGMESProfile", "CimConstants"]


def resolve_headers(path: Path, version: str) -> None:  # NOSONAR
    return
