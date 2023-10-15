import axios from 'axios';
import { GetEaWeeklyReportRequest, GetEaWeeklyReportResponse } from '@pm/models';

export const getEaWeeklyReportAPI = async (
  req: GetEaWeeklyReportRequest,
): Promise<GetEaWeeklyReportResponse> => {
  const response = await axios.get(`effort-allocation/weekly`, {
    params: req,
  });
  return response.data;
};

export const uploadEaWeeklyReportAPI = async (file: File): Promise<void> => {
  const formData = new FormData();
  formData.append('file', file);

  await axios.post(`effort-allocation/weekly`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
};
