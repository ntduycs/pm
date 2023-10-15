import { StyledEaWeeklyReportImportModal } from './styles.ts';
import { Button, notification, Upload, UploadProps } from 'antd';
import { UiConstant } from '@pm/common/constants';
import { color } from '@pm/styles';
import { UploadOutlined } from '@ant-design/icons';

export const EaWeeklyReportImportModal = () => {
  const [api, contextHolder] = notification.useNotification();

  const props: UploadProps = {
    name: 'file',
    accept: '.xlsx, .xls, .csv',
    beforeUpload: () => false,
    customRequest: async (options) => {
      const { file, onError, onSuccess } = options;
    },
    maxCount: 1,
    showUploadList: false,
    multiple: false,
    onChange(info) {
      if (info.file.status !== 'uploading') {
        console.log(info.file, info.fileList);
      }

      if (info.file.status === 'done') {
        api.success({
          message: UiConstant.SuccessMessage,
          description: 'Upload weekly report successfully',
        });
      } else if (info.file.status) {
        api.error({
          message: UiConstant.ErrorMessage,
          description: 'Upload weekly report failed',
        });
      }
    },
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
