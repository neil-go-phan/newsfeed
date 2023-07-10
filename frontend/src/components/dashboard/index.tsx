import React from 'react'
import NewsUserFeeds from './newsUserFeeds'
import RecommendedFeeds from './recommendedFeeds'
import TredingArticles from './tredingArticles'

function Dashboard() {
  return (
    <div className='dashboard'>
      <div className="dashboard__content row">
        <div className="columnLeft col-12 col-md-9">
          <div className="newsInUserFeeds">
            <NewsUserFeeds />
          </div>
          <div className="recommendedFeeds">
            <RecommendedFeeds />
          </div>
        </div>
        <div className="columnRight col-12 col-md-3">
          <div className="treding">
            <TredingArticles />
          </div>
        </div>
      </div>

    </div>
  )
}

export default Dashboard