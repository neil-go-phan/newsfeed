import React from 'react'
import NewsUserFeeds from './newsUserFeeds'
import RecommendedFeeds from './recommendedFeeds'
import TredingArticles from './tredingArticles'

function Dashboard() {
  return (
    <div className='dashboard'>
      <div className="dashboard__content">
        <div className="columnLeft">
          <div className="newsInUserFeeds">
            <NewsUserFeeds />
          </div>
          <div className="recommendedFeeds">
            <RecommendedFeeds />
          </div>
        </div>
        <div className="columnRight">
          <div className="treding">
            <TredingArticles />
          </div>
        </div>
      </div>

    </div>
  )
}

export default Dashboard