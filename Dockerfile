FROM node:14.19.3 AS ui-build
WORKDIR /usr/src/app
COPY my-app/ ./my-app/
RUN cd my-app && npm install && npm run build && npm i accounting && npm install @material-ui/core && npm i @material-ui/icons
