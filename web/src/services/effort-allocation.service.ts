import {
  GetEaWeeklyReportRequest,
  GetEaWeeklyReportResponse,
  UploadEaWeeklyReportResponse,
} from '@pm/models';
import { axios } from '@pm/services/index.ts';

export const getEaWeeklyReportAPI = async (
  req: GetEaWeeklyReportRequest,
): Promise<GetEaWeeklyReportResponse> => {
  const response = await axios.get(`effort-allocation/weekly`, {
    params: req,
  });
  return response.data;
};

export const uploadEaWeeklyReportAPI = async (
  file: File,
): Promise<UploadEaWeeklyReportResponse> => {
  const formData = new FormData();
  formData.append('file', file);

  const response = await axios.post(`effort-allocation/weekly`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });

  if (response.status > 299) {
    throw new Error(response.data.message);
  }

  return response.data;
};
