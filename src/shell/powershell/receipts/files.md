# Search Files

```powershell
# search for files
# current folder (no -recursive)
# only *.txt
# lastwritetime later than '2021-10-25'
gci ./* -include *.txt | \
? { $_.lastwritetime -gt '2021-10-25 00:00:00' } | \
select -exp name


# current folder
# lastwritetime earlier than one day ago
gi ./* -exclude *.txt | \
? { $_.lastwritetime -lt $(get-date).adddays(-1) }

```