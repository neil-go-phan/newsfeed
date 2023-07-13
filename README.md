# Newsfeed
## Description
- Newsfeed is a reliable and diverse news aggregation website. It offers a user-friendly interface and covers various categories such as politics, business, science, technology, entertainment, sports, and health. Users can access daily updates, in-depth analysis, and expert opinions. Newsfeed ensures accuracy throughautomated fact-checking. 
## Admin
- To access: 
  - go to `/admin`
- Superadmin account: 
  - username: `superadmin`
  - password: `12345678`
## Tech stack
- Frontend: 
  - Nextjs
  - Bootstrapt
  - Code convention: eslint
- Backend: 
  - Framework: Gin
  - Docker
  - Crawler: goquery
  - Message between services: gRPC
  - Database: postgres
  - Code convention: golangci-lint
  - Log: Sentry
  - Realtime push notification: Firebase cloud messages
- Deploy: 
  - Frontend: Vercel
    - Link: https://newsfeed-taupe.vercel.app/
  - Backend: 
    - Digital Ocean droplets
    - Docker
## How to run
- Device dev must install: 
  - docker
  - golang version >1.20
  - nodejs version >18.15
- Add file .env following file .env.example
- Config your firebase project
  - Change to your VAPID_KEY in file .env frontend
  - Generate your google credentials file, named it as 'newfeed_firebase.json' and copy it to '/backend/services/server' 
- Frontend: 
  - Open terminal
  - CD to frontend folder
  - Run commands:
    - `npm insall`
    - `npm run dev`
- Backend: 
  - Open terminal
  - CD to backend folder
  - Run commands:
    - `docker compose up`
## ERD
[![](https://mermaid.ink/img/pako:eNrNVtuO2jAQ_ZXIz4C4LJfNG6VslXZLVoFVpWol5I0HsOrYkeOUUuDf1yQhhDhLV-oD5CGQGV_OnDMezxb5ggCyEcjPFC8lDl64pZ_n6dibWtv04_BEIClmFiXW07eTVdEAIoWD0PIlYAVkjpVV5Y5DcslNgEHmLvgjJSlfWrHem-MADEeIo2gtJDEcEGDKDGuEmTKMUjCYH1a3HrK49umP5z6Or8KAATFBZ1gJRL6koaKCn-F-GnvfnenUcSfX0c8UgyuqNoY5ALUS5N9hFcAnWukISkpN3Jnz4IyGs6vFXO3Xk002fMGVZsQAeUjyitiG3swZ6UScT91nbzS-EUm1oAw-KF3qY5T_MowL0OtXemiAl4UNqOZMLkFaC8GYWAMxSFAipH4FfyNv-OPxSqXMLEAilj4kEec4P0SHL_GazdUmNEnHUlFfHwtCf7_rq5Yrn3lBtuOYSlRHJ471SZaRGa6_AhIfdy4l9E0I8r95fNokjF8ZjVZgFjSTnTTajL1onmVFkrtWKXndyVf30-3cQlVT9FuWp5y8wEnJl59lDut5ToIv4oqimOR9UhfPeRnOxl9czxnfGjMZvpn75IyujS1jUM9bCrkplsZL0NO-b7er13e7rAWykzsLU17si_aHEduzZqM0rrDStnRFl0bmab5Ptz3WbFvnig6Q--AuStU8BWhcjrb1Ckzw5UxY71yfGZ6jPUdycfwuU_S0_JnOOe48LQs4UA0FIHU_SnSTnWTEC1Ir0LQjW_8lsMCxbkvRC9_robpYiOmG-8heYBZBDaXKZ615bg0x_ymE_lYyTj-RvUV_kF3v9HuNu2ar1xq0OoNBd9CpoY02D1rNRrPTvu92Or12u9Xr72vob7JCu9G96_abbf26a_fv-4Pu_g1OR7WP?type=png)](https://mermaid.live/edit#pako:eNrNVtuO2jAQ_ZXIz4C4LJfNG6VslXZLVoFVpWol5I0HsOrYkeOUUuDf1yQhhDhLV-oD5CGQGV_OnDMezxb5ggCyEcjPFC8lDl64pZ_n6dibWtv04_BEIClmFiXW07eTVdEAIoWD0PIlYAVkjpVV5Y5DcslNgEHmLvgjJSlfWrHem-MADEeIo2gtJDEcEGDKDGuEmTKMUjCYH1a3HrK49umP5z6Or8KAATFBZ1gJRL6koaKCn-F-GnvfnenUcSfX0c8UgyuqNoY5ALUS5N9hFcAnWukISkpN3Jnz4IyGs6vFXO3Xk002fMGVZsQAeUjyitiG3swZ6UScT91nbzS-EUm1oAw-KF3qY5T_MowL0OtXemiAl4UNqOZMLkFaC8GYWAMxSFAipH4FfyNv-OPxSqXMLEAilj4kEec4P0SHL_GazdUmNEnHUlFfHwtCf7_rq5Yrn3lBtuOYSlRHJ471SZaRGa6_AhIfdy4l9E0I8r95fNokjF8ZjVZgFjSTnTTajL1onmVFkrtWKXndyVf30-3cQlVT9FuWp5y8wEnJl59lDut5ToIv4oqimOR9UhfPeRnOxl9czxnfGjMZvpn75IyujS1jUM9bCrkplsZL0NO-b7er13e7rAWykzsLU17si_aHEduzZqM0rrDStnRFl0bmab5Ptz3WbFvnig6Q--AuStU8BWhcjrb1Ckzw5UxY71yfGZ6jPUdycfwuU_S0_JnOOe48LQs4UA0FIHU_SnSTnWTEC1Ir0LQjW_8lsMCxbkvRC9_robpYiOmG-8heYBZBDaXKZ615bg0x_ymE_lYyTj-RvUV_kF3v9HuNu2ar1xq0OoNBd9CpoY02D1rNRrPTvu92Or12u9Xr72vob7JCu9G96_abbf26a_fv-4Pu_g1OR7WP)
