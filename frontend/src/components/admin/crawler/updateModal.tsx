import React, { useEffect, useMemo, useState } from 'react';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button, Form, InputGroup } from 'react-bootstrap';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';

type EditScheduleYupValidateProp = {
  schedule: string;
};

type Props = {
  schedule: string;
  id: number;
  handleEdit: () => void;
};

const SCHEDULE_TYPE_EVERY = '@every';
const SCHEDULE_TYPE_CRONTAB = 'crontab';
const DEFAUL_CRON_TAB = '10 * * * *';
const FAIL_MESSAGE = 'fail ';

const EditScheduleModal: React.FC<Props> = (props: Props) => {
  const [scheduleType, setScheduleType] = useState<string>('');
  const [hour, setHour] = useState<number>(0);
  const [minute, setMinute] = useState<number>(10);

  const [crontabMin, setCrontabMin] = useState<number>(10);
  const [crontabHour, setCrontabHour] = useState<number>(-1);
  const [crontabDayOfMonth, setCrontabDayOfMonth] = useState<number>(-1);
  const [crontabMonth, setCrontabMonth] = useState<number>(-1);
  const [crontabDayOfWeek, setCrontabDayOfWeek] = useState<number>(-1);

  const [errorMessage, setErrorMessage] = useState<string>('');
  const schema = yup.object().shape({
    schedule: yup.string().required('schedule must not be empty'),
  });

  const handleChooseType = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setScheduleType(e.target.value);
    if (e.target.value === SCHEDULE_TYPE_EVERY) {
      setValue('schedule', `${SCHEDULE_TYPE_EVERY} ${hour}h${minute}m`);
    }
    if (e.target.value === SCHEDULE_TYPE_CRONTAB) {
      setValue('schedule', DEFAUL_CRON_TAB);
    }
  };

  const handleUpdateResultTypeEvery = () => {
    if (scheduleType === SCHEDULE_TYPE_EVERY) {
      if (hour < 0 || minute < 0) {
        setErrorMessage('input invalid');
      } else {
        setErrorMessage('');
        setValue('schedule', `${SCHEDULE_TYPE_EVERY} ${hour}h${minute}m`);
      }
    }
  };

  const handleUpdateCrontabResult = () => {
    if (scheduleType === SCHEDULE_TYPE_CRONTAB) {
      let newCrontab = '';
      if (crontabMin >= 0 && crontabMin <= 59) {
        newCrontab += `${crontabMin} `;
      } else {
        newCrontab += '* ';
      }
      if (crontabHour >= 0 && crontabHour <= 23) {
        newCrontab += `${crontabHour} `;
      } else {
        newCrontab += '* ';
      }
      if (crontabDayOfMonth >= 1 && crontabDayOfMonth <= 31) {
        newCrontab += `${crontabDayOfMonth} `;
      } else {
        newCrontab += '* ';
      }
      if (crontabMonth >= 1 && crontabMonth <= 12) {
        newCrontab += `${crontabMonth} `;
      } else {
        newCrontab += '* ';
      }
      if (crontabDayOfWeek >= 0 && crontabDayOfWeek <= 6) {
        newCrontab += crontabDayOfWeek;
      } else {
        newCrontab += '*';
      }
      setValue('schedule', newCrontab);
    }
  };

  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm<EditScheduleYupValidateProp>({
    resolver: yupResolver(schema),
  });
  const onSubmit: SubmitHandler<EditScheduleYupValidateProp> = async (data) => {
    requestUpdateSchedule(props.id, data.schedule)
  };

  const requestUpdateSchedule = async (id: number, schedule: string) => {
    try {
      const { data } = await axiosProtectedAPI.post(
        '/crawler/update/schedule',
        {
          id: id,
          schedule: schedule,
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw FAIL_MESSAGE;
      }
      props.handleEdit();
    } catch (error: any) {
      setErrorMessage(error);
    }
  };

  useEffect(() => {
    handleUpdateResultTypeEvery();
  }, [hour, minute]);

  useEffect(() => {
    handleUpdateCrontabResult();
  }, [
    crontabMin,
    crontabHour,
    crontabMonth,
    crontabDayOfMonth,
    crontabDayOfWeek,
  ]);

  return (
    <div className="adminCrawler__modal">
      <form onSubmit={handleSubmit(onSubmit)}>
        <h2 className="adminCrawler__modal--title">Edit schedule</h2>
        <div className="adminCrawler__modal--line" />

        <div className="field mb-3">
          <label> Schedule type </label>
          <Form.Select
            onChange={(e) => handleChooseType(e)}
            aria-label="Choose type"
          >
            <option value="Choose type">Choose type</option>
            <option value={SCHEDULE_TYPE_EVERY}>{SCHEDULE_TYPE_EVERY}</option>
            <option value={SCHEDULE_TYPE_CRONTAB}>
              {SCHEDULE_TYPE_CRONTAB}
            </option>
          </Form.Select>
        </div>

        {scheduleType === SCHEDULE_TYPE_EVERY ? (
          <div className="field mb-3">
            <div className="row">
              <div className="hour col-6">
                <label>Hour</label>
                <Form.Control
                  placeholder="Please type hour"
                  type="number"
                  required
                  className="w-50"
                  value={hour}
                  onChange={(e) => setHour(+e.target.value)}
                />
              </div>
              <div className="minute col-6">
                <label>Minute</label>
                <Form.Control
                  placeholder="Please type minute"
                  type="number"
                  required
                  className="w-50"
                  value={minute}
                  onChange={(e) => setMinute(+e.target.value)}
                />
              </div>
            </div>
          </div>
        ) : (
          <div className="field mb-3">
            <div className="row">
              <div className="col-4">
                <label>Minute (0-59)</label>
                <Form.Control
                  placeholder="Please type minute"
                  type="number"
                  required
                  className="w-100"
                  value={crontabMin}
                  onChange={(e) => setCrontabMin(+e.target.value)}
                />
              </div>
              <div className="col-4">
                <label>Hour (0-23)</label>
                <Form.Control
                  placeholder="Please type hour"
                  type="number"
                  required
                  className="w-100"
                  value={crontabHour}
                  onChange={(e) => setCrontabHour(+e.target.value)}
                />
              </div>
              <div className="col-4">
                <label>Day of month (1-31)</label>
                <Form.Control
                  placeholder="Please type day of month"
                  type="number"
                  required
                  className="w-100"
                  value={crontabDayOfMonth}
                  onChange={(e) => setCrontabDayOfMonth(+e.target.value)}
                />
              </div>
            </div>
            <div className="row my-3">
              <div className="col-4">
                <label>Month (1 - 12)</label>
                <Form.Control
                  placeholder="Please type month"
                  type="number"
                  required
                  className="w-100"
                  value={crontabMonth}
                  onChange={(e) => setCrontabMonth(+e.target.value)}
                />
              </div>
              <div className="col-4">
                <label>Day of week (0 - 6) (Sunday=0)</label>
                <Form.Control
                  placeholder="Please type day of week"
                  type="number"
                  required
                  className="w-100"
                  value={crontabDayOfWeek}
                  onChange={(e) => setCrontabDayOfWeek(+e.target.value)}
                />
              </div>
            </div>
          </div>
        )}

        <div className="field">
          <label> Result </label>
          <InputGroup className="mb-3">
            <Form.Control
              {...register('schedule')}
              placeholder="Please choose schedule type"
              type="text"
              required
              disabled
            />
          </InputGroup>

          {errors.schedule && (
            <p className="errorMessage">{errors.schedule.message}</p>
          )}
        </div>
        {errorMessage !== '' && <p className="errorMessage">{errorMessage}</p>}
        <Button className="px-4 m-3" variant="primary" type="submit">
          Edit schedule
        </Button>
      </form>
    </div>
  );
};

export default EditScheduleModal;
