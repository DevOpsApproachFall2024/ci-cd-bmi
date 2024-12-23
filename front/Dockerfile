FROM node:22 AS dependencies
WORKDIR /client
COPY package.json package-lock.json /app/
RUN npm clean-install

FROM node:22 AS base
WORKDIR /client
COPY --from=dependencies /app/node_modules ./node_modules
COPY . .

FROM base AS development
CMD ["npm", "run", "dev"]

FROM base AS build
RUN npm run build

FROM build AS production
COPY --from=base /client/node_modules ./node_modules
COPY --from=build /client/dist ./dist
RUN npm install -g serve
EXPOSE 3000
CMD ["serve", "-s", "dist", "-l", "3000"]
