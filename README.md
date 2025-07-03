# FlipOS+
This will be the start of a fixed up stock OS experience. The idea is to cleanup the stock OS to make it usable.

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

## Order systems

### Current features
Systems will be ordered based on the contents of order.txt.
The order.txt file will be generated the first time you run the app.
Modify this file and put the systems in the order you want.
Run it again and it will order them based on your preference.

### Future development

#### GUI
A graphical user interface is needed to remove the need of using the order.txt file.

### Stock Behavior

#### Ordering
The miyoo main ui interface orders the systems alphabetically. So by adding spaces around the name of the system we can force the order we want.

## Syncthing
The syncthing app after running will run in the background on port 8384.

### Future development
- Syncthing app should be able to be run at boot if configured to do so.
- Syncthing should use a different account than root.
- Autoconfigure the Rom and Save directories.

# Credits
- [Christopher Cromer](https://github.com/cromerc) (Apps)
