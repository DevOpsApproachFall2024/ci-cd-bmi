FROM node:22 AS base
WORKDIR /client
COPY package.json package-lock.json ./
RUN npm clean-install

FROM base AS development
CMD ["npm", "run", "dev"]

FROM base AS build
RUN npm run build

FROM build AS production
COPY --from=base /client/node_modules ./node_modules
COPY --from=build /client/dist ./dist
RUN npm install -g serve
EXPOSE 5000
CMD ["serve", "-s", "dist", "-l", "3000"]