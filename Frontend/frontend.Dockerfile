FROM node:14

WORKDIR /api-frontend

COPY package*.json ./

RUN npm install && npm i accounting && npm install @material-ui/core && npm i @material-ui/icons

EXPOSE 3000

CMD ["npm", "start"]
