# go_lang_packager
packager for my new language written in go

## Key Files
- tfpackage.json
  - This file contains information about the current overall package
- lang_info.json
  - This file contains information about the language being created
  - It also contains information about required .json files(that the packager can understand)
  - Example: lang_info.json is a required module by default.
  - You can require it inside your TypeFast file and use it during compilation(example: having compilation information on what files is running)
