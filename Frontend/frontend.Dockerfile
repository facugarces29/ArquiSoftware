FROM node:14.19-buster

WORKDIR /api-frontend

COPY package*.json ./

RUN npm install 

EXPOSE 3000

CMD ["npm", "start"]
