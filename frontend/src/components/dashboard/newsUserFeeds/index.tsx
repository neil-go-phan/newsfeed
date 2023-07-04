import React, { useEffect, useState } from 'react';
import DashboardFeeds from './feeds';
import Image from 'next/image';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';

const IMAGE_WIDTH = 290;

const REQUEST_NEWEST_UPDATED_SOURCES_FAIL_MESSAGE = 'request newest updated source fail'

function NewsUserFeeds() {
  const [sourceIDs, setSourceIDs] = useState<Array<number>>([]);
  
  useEffect(() => {
    requestNewstSourceIDs()
  }, [])

  const requestNewstSourceIDs = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('/follow/get/new-update',);
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_NEWEST_UPDATED_SOURCES_FAIL_MESSAGE;
      }

      setSourceIDs(data.articles_source_ids);
    } catch (error: any) {
      setSourceIDs([]);
    }
  };

  return (
    <div className="wrapper">
      <div className="title">What's new in your feeds</div>
      {sourceIDs.length > 0 ? (
        <div className="listFeeds">
          {sourceIDs.map((id) => (
            <DashboardFeeds sourceid={id} key={`DashboardFeeds ${id}`}/>
          ))}
        </div>
      ) : (
        <div className="nothingNew">
          <Image
            alt="nothing new"
            src={'/images/section-all-read-aqua.svg'}
            width={IMAGE_WIDTH}
            height="0"
            style={{ height: 'auto' }}
          />
          <p>You have read everything!</p>
        </div>
      )}
    </div>
  );
}

export default NewsUserFeeds;
