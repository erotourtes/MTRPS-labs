FROM node:21-alpine3.18

WORKDIR /app

COPY ./starter-project-nodejs/package.* .
RUN npm install 

COPY ./starter-project-nodejs .

CMD ["npm", "start"]
