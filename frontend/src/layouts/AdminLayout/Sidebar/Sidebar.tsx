import React, { useEffect, useState } from 'react'
import classNames from 'classnames'
import SidebarNav from './SidebarNav'

export default function Sidebar(props: { isShow: boolean; isShowMd: boolean }) {
  const { isShow, isShowMd } = props
  const [isNarrow, setIsNarrow] = useState(false)

  // On first time load only
  useEffect(() => {
    if (localStorage.getItem('isNarrow')) {
      setIsNarrow(localStorage.getItem('isNarrow') === 'true')
    }
  }, [setIsNarrow])

  return (
    <div
      className={classNames('adminSidebar d-flex flex-column position-fixed h-100', {
        'sidebar-narrow': isNarrow,
        show: isShow,
        'md-hide': !isShowMd,
      })}
      id="sidebar"
    >
      <div className="sidebar-brand d-none d-md-flex align-items-center justify-content-center">
        <h1 className="fs-6">Football news aggregation</h1>
      </div>

      <div className="sidebar-nav flex-fill">
        <SidebarNav />
      </div>
    </div>
  )
}

export const SidebarOverlay = (props: { isShowSidebar: boolean; toggleSidebar: () => void }) => {
  const { isShowSidebar, toggleSidebar } = props

  return (
    <div
      tabIndex={-1}
      aria-hidden
      className={classNames('adminSidebar-overlay position-fixed top-0 bg-dark w-100 h-100 opacity-50', {
        'd-none': !isShowSidebar,
      })}
      onClick={toggleSidebar}
    />
  )
}
