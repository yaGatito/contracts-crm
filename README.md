# How to build

### Web
*From root project folder*
```sh
cd web
npm install
npm run build
```
### Wails binary
```sh
cd web
wails build -platform linux/amd64 -clean
```

### Run
```sh
./build/bin/contracts-crm
```
