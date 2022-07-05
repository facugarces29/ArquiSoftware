FROM node:14.19-buster

WORKDIR /api-frontend

COPY . ./api-frontend

RUN npm install 

RUN npm build

CMD ["npm", "start", "build"]
