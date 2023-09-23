import { PropsWithChildren } from 'react';
import { ConfigProvider } from 'antd';

export const AntdProvider = ({ children }: PropsWithChildren) => {
  return (
    <ConfigProvider
      theme={{
        components: {
          Notification: {
            width: 500,
          },
        },
        token: {},
      }}
    >
      {children}
    </ConfigProvider>
  );
};
