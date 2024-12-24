FROM node:22 AS dependencies
WORKDIR /client/
COPY package.json package-lock.json ./
RUN npm ci


FROM node:22 AS base
WORKDIR /client
COPY --from=dependencies /client/node_modules ./node_modules
COPY . .

FROM base AS development
CMD ["npm", "run", "dev"]

FROM base AS build
RUN npm run build

FROM build AS production
COPY --from=base /client/node_modules ./node_modules
COPY --from=build /client/dist ./dist
RUN npm install -g serve
EXPOSE 5173
CMD ["serve", "-s", "dist", "-l", "5173"]
