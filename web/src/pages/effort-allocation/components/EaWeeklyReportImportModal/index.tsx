import { StyledEaWeeklyReportImportModal } from './styles.ts';
import { Button, notification, Upload, UploadProps } from 'antd';
import { UiConstant } from '@pm/common/constants';
import { color } from '@pm/styles';
import { UploadOutlined } from '@ant-design/icons';
import { useMutation } from '@tanstack/react-query';
import { uploadEaWeeklyReportAPI } from '@pm/services';
import { RcFile } from 'antd/es/upload';
import { Apis } from '@pm/common/utils';
import { useResetRecoilState, useSetRecoilState } from 'recoil';
import { effortAllocationState } from '@pm/atoms';

export const EaWeeklyReportImportModal = () => {
  const [api, contextHolder] = notification.useNotification();
  const setEaState = useSetRecoilState(effortAllocationState);
  const resetEaState = useResetRecoilState(effortAllocationState);
  const { mutateAsync: uploadWeeklyReport } = useMutation({
    mutationKey: ['effort-allocations'],
    mutationFn: uploadEaWeeklyReportAPI,
  });

  const props: UploadProps = {
    name: 'file',
    accept: '.xlsx, .xls, .csv',
    customRequest: async (options) => {
      const { file, onError, onSuccess } = options;

      try {
        const { items: eaState } = await uploadWeeklyReport(file as RcFile);
        setEaState(eaState);
        api.success({
          message: UiConstant.SuccessMessage,
          description: 'Upload weekly report successfully',
        });
      } catch (error) {
        resetEaState();
        api.error({
          message: UiConstant.ErrorMessage,
          description: Apis.getErrorDescription(error),
        });
      }
    },
    maxCount: 1,
    showUploadList: false,
    multiple: false,
    progress: {
      strokeColor: {
        '0%': color.blue,
        '100%': color.green,
      },
      strokeWidth: 3,
      format: (percent) => `${parseFloat((percent ?? 0).toFixed(2))}%`,
    },
  };

  return (
    <>
      {contextHolder}
      <StyledEaWeeklyReportImportModal>
        <Upload {...props}>
          <Button
            type='primary'
            icon={<UploadOutlined />}
          >
            Upload
          </Button>
        </Upload>
      </StyledEaWeeklyReportImportModal>
    </>
  );
};
