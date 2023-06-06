import { _ROUTES } from '@/helpers/constants'
import { setCookie } from 'cookies-next'
import { useRouter } from 'next/router'
import React, { useEffect } from 'react'

function Token() {
  const route = useRouter()
  useEffect(()=>{
    if (route.query.access_token) {
      const token = route.query.access_token
      setCookie('access_token', token) 
    }
    if (route.query.refresh_token) {
      const token = route.query.refresh_token
      setCookie('refresh_token', token)
    }
    route.push(_ROUTES.USER_HOME)
  })
  return (
    <div></div>
  )
}

export default Token