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
[![](https://mermaid.ink/img/pako:eNrNV21v2jAQ_iuRP9OqlJRCvjEKExtrqkBVaaoUuckBVhM7cpwxBvnvM-SFJA6UqZogHyC5O_vunnvxeY0c5gIyEPAHgucc-69Uk49ljgcTbZ18bJ8QOMGeRlzt6fueKogPocB-oDkcsADXxkKrY0eBe4ztggcJu6BRcELnGsU-aArVhdDhJBCE0YQXJ39PA-vHaDIZmY9nsV41FKggYqWQfRAL5p7s1jYc9nHfOPPArjiYsgLgPglDuWlRIC5GOt5cXW3Wqh5DcxgVmFAV3hOXPE8G1nlioQYjkrq3-aQwAhyGS8bVgICPiadQQ-ypiboLwC5bh2WEEwQ2W7g2KdwVjHrWdNSXdHtiPlv9wYWkrkxcDz5O0QLPI_RdIc5A7l_LIT6eFxQQKoDPgWsz5nlsCa4CgmABcbY5XEG4b_VexmdKMzU5WMQd2Hmc23kSHA7HS88Wq0AFHXNBHJleLvl1kFcfrnzlkbBlMrVWZUwcyY7FQ9VdZwFuVNKcIJ8uDO0UkCNhS2pDqQJDewOP0fmUaeU6uYg4f7Y89kqC6M0j4QLU9qOCfgRbrQSuguYm6dcZPW9BWhYM8_Gb-eVyDv26JfKXV5fsuUDdCi9vKRSWdg6awyIqFC935QdcYllJ0hSXOGngWdIacnOpljpgzlLJ3nTw1bRGg0vEMGu5nhfJV6xMGFPzadQ_t-FpIOS6OeOrYseo9Ss1PbM9j1Aeh4MNpFwUm3SHvXgiPTTHY_OlBMrn3K-MIXWzWk1lF6W2Cb3N54hK9W7NmJGVeWp7tcoVAD6Qtwa9h38B4HQPD8yqH58aiUlxuZupZ0VBbJNiU41vSeTI6bOXtscyt6z_D0isqjzJ4ZLwIbeH_R_2ozkdDUeyUpTbxDF70wmGcHjDoZw42DvQqs3q7geMQQ3ky4sJJq68fu4seEViAbK4kSFfXZjhSI7Z6JXGUlSehGyyog4yZtgLoYGSAksvrTk1wPQnY_Jb8Cj5RMYa_UZGp33d0u_1W73Z7ujN21a3gVbIuGp2utd37W5T19s39-1u6_42bqA_ux2a161Oq9tq6zd3zU6z09E78V-CPakH?type=png)](https://mermaid.live/edit#pako:eNrNV21v2jAQ_iuRP9OqlJRCvjEKExtrqkBVaaoUuckBVhM7cpwxBvnvM-SFJA6UqZogHyC5O_vunnvxeY0c5gIyEPAHgucc-69Uk49ljgcTbZ18bJ8QOMGeRlzt6fueKogPocB-oDkcsADXxkKrY0eBe4ztggcJu6BRcELnGsU-aArVhdDhJBCE0YQXJ39PA-vHaDIZmY9nsV41FKggYqWQfRAL5p7s1jYc9nHfOPPArjiYsgLgPglDuWlRIC5GOt5cXW3Wqh5DcxgVmFAV3hOXPE8G1nlioQYjkrq3-aQwAhyGS8bVgICPiadQQ-ypiboLwC5bh2WEEwQ2W7g2KdwVjHrWdNSXdHtiPlv9wYWkrkxcDz5O0QLPI_RdIc5A7l_LIT6eFxQQKoDPgWsz5nlsCa4CgmABcbY5XEG4b_VexmdKMzU5WMQd2Hmc23kSHA7HS88Wq0AFHXNBHJleLvl1kFcfrnzlkbBlMrVWZUwcyY7FQ9VdZwFuVNKcIJ8uDO0UkCNhS2pDqQJDewOP0fmUaeU6uYg4f7Y89kqC6M0j4QLU9qOCfgRbrQSuguYm6dcZPW9BWhYM8_Gb-eVyDv26JfKXV5fsuUDdCi9vKRSWdg6awyIqFC935QdcYllJ0hSXOGngWdIacnOpljpgzlLJ3nTw1bRGg0vEMGu5nhfJV6xMGFPzadQ_t-FpIOS6OeOrYseo9Ss1PbM9j1Aeh4MNpFwUm3SHvXgiPTTHY_OlBMrn3K-MIXWzWk1lF6W2Cb3N54hK9W7NmJGVeWp7tcoVAD6Qtwa9h38B4HQPD8yqH58aiUlxuZupZ0VBbJNiU41vSeTI6bOXtscyt6z_D0isqjzJ4ZLwIbeH_R_2ozkdDUeyUpTbxDF70wmGcHjDoZw42DvQqs3q7geMQQ3ky4sJJq68fu4seEViAbK4kSFfXZjhSI7Z6JXGUlSehGyyog4yZtgLoYGSAksvrTk1wPQnY_Jb8Cj5RMYa_UZGp33d0u_1W73Z7ujN21a3gVbIuGp2utd37W5T19s39-1u6_42bqA_ux2a161Oq9tq6zd3zU6z09E78V-CPakH)
