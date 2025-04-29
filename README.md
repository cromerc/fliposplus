# Miyoo OS+
## Cleanup game names
### Current features
Subdirectories will also be scanned and cleaned up.
Subdirectory names will be skipped.
#### Remove parenthesis
Parenthesis will be removed from the rom file names only.
#### Move articles
Articles will be moved to the begging of the file name.
### Future development
#### Generate cache files
If the directory is a rom directory and there is no cache file, a cache file should be created.
This should take into consideration roms on both the main sd card and the secondary sd card.
#### Cleanup recent and favorite names
The names shown in  recent and favorite games are not part of the db cache, these will need xml modification to clean up.
### Stock Behavior
#### Database cache
##### Generation
The miyoo stock os generates roms.db files when you look at the games inside of a directory.
##### Updating
The miyoo stock os does not seem to update the list if the cache exists.
Is there a timeout before the stock os rescans the directory?
#### Refresh roms
The refresh roms option in the miyoo stock os deletes the cache db files.
If this happens
# Credits
- [Christopher Cromer](https://github.com/cromerc) (Apps)
