# Portal Frontend Next project

First, run the development server:

```bash
yarn dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.


## To generate a static distribution

First make sure that  `next.config.js` has: 
```js
output: 'export', 
```

Run

```shell
npm run build
```

Copy the contents of the `out/` directory to `../portal/kodata`