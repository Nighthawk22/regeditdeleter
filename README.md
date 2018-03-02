# Regedit deleter

The prupose of this little programm is to delete Windows registry keys as it seems 
that normal bat scripts are not working at Windows 10 anymore.

## Usage

The idea is that this application will be run within a bat script as Administrator.

Call:
```
regeditdeleter.exe --root LOCAL_MACHINE --path SYSTEM\CurrentControlSet\Control\Lsa\LmCompatibilityLevel
```

## Parameters

| Parameter  | Purpose                                         |
| ---------- | ------------------------------------------------|
| root       | The root of the registry. E.g.: `LOCAL_MACHINE` |
| path       | The path to the key to be deleted without the root. E.g.: `SYSTEM\CurrentControlSet\Control\Lsa\LmCompatibilityLevel` |

### Possible roots

* CLASSES_ROOT
* CURRENT_USER
* LOCAL_MACHINE
* USERS
* CURRENT_CONFIG
* PERFORMANCE_DATA

