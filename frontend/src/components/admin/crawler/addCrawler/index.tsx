import React, { useState } from 'react'
import Tutorial from './tutorial'
import InputUrl from './inputUrl'
import TestResult from './testResult'

function AddCrawler() {
  const [url, setUrl] = useState<string>('')
  const handleInputUrl= (url:string) => {
    setUrl(url)
  }
  return (
    <div className='addCrawler'>
      <Tutorial />
      <InputUrl handleInputUrl={handleInputUrl}/>
      {url === '' ? <></>: <TestResult url={url}/>}
    </div>
  )
}

export default AddCrawler