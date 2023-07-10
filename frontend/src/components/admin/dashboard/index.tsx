import React, { forwardRef, useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import BarChartOnHour from './barChartOnHour';
import DatePicker from 'react-datepicker';
import BarChartOnDay from './barChartOnDay';
import { ThreeDots } from 'react-loader-spinner';
import { Button } from 'react-bootstrap';
import { toastifyError } from '@/helpers/toastify';

export type ChartOnDayData = {
  time: number;
  amount_of_jobs: number;
  cronjobs: Array<CronjobDayBar>;
};

export type CronjobDayBar = {
  name: string;
  times: number;
};

export type ChartOnHourData = {
  minute: number;
  amount_of_jobs: number;
  cronjobs: Array<CronjobBar>;
};

export type CronjobBar = {
  name: string;
  start_at: string;
  end_at: string;
  new_articles_count: number;
};

type CustomInputProps = {
  onClick: React.MouseEventHandler<HTMLButtonElement>;
  date: Date | undefined;
};

function AdminDashboard() {
  const [chartOnHourData, setChartOnHourData] = useState<Array<ChartOnHourData>>();
  const [chartOnDayData, setChartOnDayData] = useState<Array<ChartOnDayData>>();
  const [choosenDay, setChoosenDay] = useState<Date>();
  const [choosenHour, setChoosenHour] = useState<Date>()
  const router = useRouter();

  const formatTimeHour = (date: Date | undefined): string => {
    if (date) {
      let month = `${date.getMonth() + 1}`;
      if (date.getMonth() + 1 < 10) {
        month = `0${date.getMonth() + 1}`;
      }
      let day = `${date.getDate()}`;
      if (date.getDate() < 10) {
        day = `0${date.getDate()}`;
      }
      return `${date.getFullYear()}-${month}-${day} ${date.getHours()}`;
    }
    return '';
  };

  const formatTimeDay = (date: Date | undefined): string => {
    if (date) {
      let month = `${date.getMonth() + 1}`;
      if (date.getMonth() + 1 < 10) {
        month = `0${date.getMonth() + 1}`;
      }
      let day = `${date.getDate()}`;
      if (date.getDate() < 10) {
        day = `0${date.getDate()}`;
      }
      return `${date.getFullYear()}-${month}-${day}`;
    }
    return '';
  };

  const handleChooseHour = (hour: number) => {
    if (choosenDay) {
      const temp = new Date(choosenDay)
      temp.setHours(hour)
      setChoosenHour(temp)
    }
  }

  const requestChartByHour = async (timeString: string) => {
    try {
      const { data } = await axiosProtectedAPI.get('/crawler/cronjob/hour', {
        params: { time: timeString },
      });
      setChartOnHourData(data.cronjobs);
    } catch (error) {
      toastifyError('Error occurred while get list cronjob')
    }
  };

  const requestChartByDay = async (timeString: string) => {
    try {
      const { data } = await axiosProtectedAPI.get('/crawler/cronjob/day', {
        params: { time: timeString },
      });
      setChartOnDayData(data.cronjobs);
    } catch (error) {
      toastifyError('Error occurred while get chart data')
    }
  };

  useEffect(() => {
    const date = new Date();
    setChoosenDay(date);
    setChoosenHour(date)
  }, [router.asPath]);
  useEffect(() => {
    if (choosenHour) {
      requestChartByHour(formatTimeHour(choosenHour))
    }
  }, [choosenHour]);
  useEffect(() => {
    if (choosenDay) {
      requestChartByDay(formatTimeDay(choosenDay));
      requestChartByHour(formatTimeHour(choosenDay))
    }
  }, [choosenDay]);
  const handleOnClickChooseDay = (dateChose: Date) => {
    setChoosenDay(dateChose);
    requestChartByDay(formatTimeDay(dateChose));
  };
  const BtnCustomInput = forwardRef<
    HTMLButtonElement,
    {
      onClick: React.MouseEventHandler<HTMLButtonElement>;
      date: Date | undefined;
    }
  >(({ onClick, date }, ref) => (
    <Button className="btnTriggerDate bg-success" onClick={onClick} ref={ref}>
      {date ? formatTimeDay(date) : 'Choose date'}
    </Button>
  ));

  const CustomInput = React.forwardRef<HTMLButtonElement, CustomInputProps>(
    ({ onClick, date }, ref) => (
      <BtnCustomInput onClick={onClick} ref={ref} date={date} />
    )
  );
  return (
    <div className="adminDashboard">
      <div className="adminDashboard__title">Dashboard</div>
      <div className="adminDashboard__chooseDate">
        <div className="adminDashboard__chooseDate--title mx-3">
          Choose date
        </div>
        <DatePicker
          selected={choosenDay}
          onChange={(date: Date) => handleOnClickChooseDay(date)}
          customInput={<CustomInput onClick={() => {}} date={choosenDay} />}
        />
      </div>
      <div className="adminDashboard__dayChart">
        <div className="adminDashboard__dayChart--title">Cronjob run in day</div>
        {chartOnDayData ? (
          <div className="adminDashboard__dayChart--warper">
            <BarChartOnDay
              title={formatTimeDay(choosenDay)}
              chartData={chartOnDayData}
              handleChooseHour={handleChooseHour}
            />
          </div>
        ) : (
          <div className="adminDashboard__dayChart--loading">
            <ThreeDots
              height="50"
              width="50"
              radius="9"
              color="#4fa94d"
              ariaLabel="three-dots-loading"
              visible={true}
            />
          </div>
        )}
      </div>
      <div className="adminDashboard__hourChart">
        <div className="adminDashboard__hourChart--title">Cronjob run in hour</div>
        {chartOnHourData && choosenHour ? (
          <div className="adminDashboard__dayChart--warper">
            <BarChartOnHour
              title={formatTimeHour(choosenHour)}
              hour={choosenHour.getHours()}
              chartData={chartOnHourData}
            />
          </div>
        ) : (
          <div className="adminDashboard__dayChart--loading">
            <ThreeDots
              height="50"
              width="50"
              radius="9"
              color="#4fa94d"
              ariaLabel="three-dots-loading"
              visible={true}
            />
          </div>
        )}
      </div>
    </div>
  );
}

export default AdminDashboard;
