import Image from 'next/image';
import React, { useContext } from 'react';
import { compareItems } from '../pricing/compare';
import FeatureCel from './featureCel';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { alertError, alertSuccess } from '@/helpers/alert';
import { setCookie } from 'cookies-next';
import { RoleContext } from '@/common/contexts/roleContext';

const ICON_SIZE = 60;

const UPDATE_USER_ROLE_FAIL_MESSAGE = 'fali';
const UPDATE_USER_ROLE_SUCCESS_MESSAGE = 'You have become a premium user';
const GET_ROLE_FAIL_MESSAGE = 'fail';

function UserPlan() {
  const handleUpgrateToPremium = () => {
    requestUpgrateToPremium();
  };

  const { setRole } = useContext(RoleContext);

  const requestUpgrateToPremium = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('/auth/update/premium');
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw UPDATE_USER_ROLE_FAIL_MESSAGE;
      }
      setCookie('access_token', data.access_token);
      setCookie('refresh_token', data.refresh_token);
      alertSuccess(UPDATE_USER_ROLE_SUCCESS_MESSAGE);
      requestGetRole();
    } catch (error: any) {
      alertError(error);
    }
  };
  const requestGetRole = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('role/get');
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw GET_ROLE_FAIL_MESSAGE;
      }
      setRole(data.role);
    } catch (error: any) {
      alertError(
        'fail to get role, please logout and login again to update your role'
      );
    }
  };

  return (
    <div className="plan">
      <div className="plan__header">Upgrade for even more smart features</div>
      <div className="plan__table ">
        <div className="table-header">
          <div className="d-flex">
            <div className="table-cel info col-6">
              <div className="heading">Try Premium plan for free!</div>
              <div className="description">
                Our free trial starts you off on the Premium plan. When you're
                ready to buy, you'll choose the plan that's right for you.
              </div>
              <div className="upgradeBtn">
                <button onClick={handleUpgrateToPremium}>
                  Start free trial
                </button>
              </div>
            </div>
            <div className="table-cel freePlan col-3">
              <div className="icon">
                <Image
                  alt="basic tier"
                  src={'/images/pricing-free-plan-aqua.svg'}
                  width={ICON_SIZE}
                  height={ICON_SIZE}
                />
              </div>

              <div className="name">Basic</div>

              <div className="price">Free</div>
              <div className="introduce">Our free serivces</div>
            </div>
            <div className="table-cel premiumPlan col-3">
              <div className="icon">
                <Image
                  alt="basic tier"
                  src={'/images/pricing-pro-plan-aqua.svg'}
                  width={ICON_SIZE}
                  height={ICON_SIZE}
                />
              </div>

              <div className="name">Premium</div>

              <div className="price">
                $1 <span>/month</span>
              </div>
              <div className="introduce">More features unlock</div>
              <div className="upTierBtn">
                <button onClick={handleUpgrateToPremium}>Upgrade</button>
              </div>
            </div>
          </div>
        </div>
        <div className="table-row">
          {compareItems.map((item) => (
            <FeatureCel key={`feature cel ${item.description}`} item={item} />
          ))}
        </div>
      </div>
    </div>
  );
}

export default UserPlan;
