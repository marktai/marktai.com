FROM nginx:1.21.6-alpine

RUN rm /etc/nginx/conf.d/default.conf

ADD ./dhparams.pem /etc/ssl/dhparams.pem
ADD ./nginx.conf /etc/nginx/
ADD ./conf.d/ /etc/nginx/conf.d/
ADD ./letsencrypt/ /etc/letsencrypt
