import React, { useState } from 'react'
import FeedsHeader from './header'
import FeedsContent from './content'

function FeedsComponent() {
  const [isOpenSidebar, setIsOpenSidebar] = useState<boolean>(true)
  const handleToggleSidebar = () => {
    setIsOpenSidebar(!isOpenSidebar)
  }
  return (
    <div className='feeds'>
      <FeedsHeader isOpenSidebar={isOpenSidebar} handleToggleSidebar={handleToggleSidebar}/>
      <FeedsContent isOpenSidebar={isOpenSidebar} />
    </div>
  )
}

export default FeedsComponent