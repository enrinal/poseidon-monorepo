FROM node:latest as builder
WORKDIR /usr/app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build


FROM node:latest
WORKDIR /usr/app
COPY package*.json ./
RUN npm install

COPY . .
COPY .env .

EXPOSE 8088
CMD ["npm", "run", "start"]