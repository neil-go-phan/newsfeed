import { _ROUTES } from '@/helpers/constants';
import { faCompass } from '@fortawesome/free-regular-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import Link from 'next/link';
import React from 'react';
import CompareFeature from './compare';

function PricingComponent() {
  return (
    <div className="pricing text-center pt-5">
      <h1 className="page-header pt-5 heading-h1 blue-bottom pb-5">Pricing</h1>
      <div className="page-subheader py-3 mt-4 text-center mx-auto mb-5 text-muted-themed prg-lg">
        Building your newsfeed costs less than you think
      </div>
      <div className="my-5">
        <div className="container-xxl py-5 px-3 px-sm-5">
          <div className="row row-cols-1 row-cols-md-1 row-cols-lg-2 flex-wrap">
            <div className="col flex-column mb-4 mb-sm-3 mb-md-3">
              <div className="card py-5 px-4 py-md-0 p-lg-5 flex flex-column flex-md-row flex-lg-column">
                <div className="d-flex flex-column align-items-md-start align-items-lg-center col-md-4 col-lg-12">
                  <h5 className="card-title heading-h5">Basic</h5>
                  <div className="card-content minH px-lg-2 py-3 prg-sm text-muted-themed">
                    Great for beginners. Subscribe to RSS feeds, search in your
                    articles, own your personal read later list.
                  </div>
                </div>
                <div className="col-1 col-lg-12"></div>
                <div className="mt-lg-auto flex flex-column flex-md-row flex-lg-column col-12 col-md-7 col-lg-12 justify-content-between">
                  <div className="col-12 col-md-4 col-lg-12">
                    <div className="card-text">
                      <strong>Free forever</strong>
                    </div>
                    <p className="card-text text-muted py-2 prg-sm text-muted-themed">
                      No credit card required
                    </p>
                  </div>
                  <div className="col-12 col-md-1 col-lg-12"></div>
                  <Link
                    className="btn btn-tertiary btn-lg d-inline-block"
                    href={`${_ROUTES.LOGIN_PAGE}?redirectTo=${_ROUTES.FEEDS_PLAN}`}
                  >
                    Create account
                  </Link>
                </div>
              </div>
            </div>
            <div className="col flex-column mb-4 mb-sm-3 mb-md-3">
              <div className="card py-5 px-4 py-md-0 p-lg-5 flex flex-column flex-md-row flex-lg-column">
                <div className="d-md-none d-lg-block label-color position-absolute w-100">
                  <span className="label-accent">MOST POPULAR</span>
                </div>
                <div className="d-flex flex-column align-items-md-start align-items-lg-center col-md-4 col-lg-12">
                  <h5 className="card-title heading-h5">Permium</h5>
                  <div className="card-content minH px-lg-2 py-3 prg-sm text-muted-themed">
                    Add your own new source. Follow social media feeds and
                    websites without RSS.
                  </div>
                  <div className="col-1 col-lg-12"></div>
                  <div className="mt-lg-auto flex flex-column flex-md-row flex-lg-column col-12 col-md-7 col-lg-12 justify-content-between">
                    <div className="col-12 col-md-4 col-lg-12">
                      <div className="card-text">
                        <strong>$2</strong>/month
                      </div>
                      <p className="card-text text-muted py-2 prg-sm text-muted-themed">
                        Billed annually, or $1 billed monthly
                      </p>
                    </div>
                    <div className="col-12 col-md-1 col-lg-12"></div>
                    <Link
                      className="btn btn-accent btn-lg d-inline-block"
                      href={`${_ROUTES.LOGIN_PAGE}?redirectTo=${_ROUTES.FEEDS_PLAN}`}
                    >
                      Upgrade to Pro
                    </Link>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <CompareFeature />
      </div>
    </div>
  );
}

export default PricingComponent;
