FROM node:14.19.3

WORKDIR /usr/src/app/mi-app

COPY package*.json ./

RUN npm install

EXPOSE 3000

CMD ["npm", "start"]
