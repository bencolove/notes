# Manage All Files on Storage Device
1. media files (`FileProvider` and `MediaStore`)
1. non media files(`FileProvider`)
1. all other files

## Access Files Siliently (no user interactiion)
by:
1. declare: *`MANAGE_EXTERNAL_STORAGE`* in *manifest*
1. intent: *`ACTION_MANAGE_ALL_FILES_ACCESS_PERMISSION`* to ask user to grant permissions: **Allow access to manage all files**

verify:
`Environment.isExternalStorageManager()`

## Allowed Operations
* *`shared storage`*
* `MediaStore.Files` table
* USB and SD card
* internal??? storage except for `/sdcard/Andoird/`, `/Android/data/`