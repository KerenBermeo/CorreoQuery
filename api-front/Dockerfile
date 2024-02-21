FROM node:21.2.0-alpine

RUN  npm install -g http-server

RUN mkdir /frontend

WORKDIR /frontend

COPY ./package*.json ./

RUN npm install

COPY . .

# RUN npm run build

EXPOSE 5173

CMD ["http-server", "dist"]