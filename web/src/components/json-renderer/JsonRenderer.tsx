import { StyledJsonRenderer } from '@pm/components/json-renderer/styles.ts';
import JSONPretty from 'react-json-pretty';
import React from 'react';

type JsonRendererProps = {
  json: object;
};

const JsonRenderer = ({ json }: JsonRendererProps) => {
  return (
    <StyledJsonRenderer>
      <JSONPretty
        data={json}
        theme={{
          main: 'line-height:1.3;color:#444444;background:inherit;overflow:auto;',
          error: 'line-height:1.3;color:#444444;background:#272822;overflow:auto;',
          key: 'color:#f92672;',
          string: 'color:#444444;',
          value: 'color:#1677ff;',
          boolean: 'color:#1677ff;',
        }}
      />
    </StyledJsonRenderer>
  );
};

export default JsonRenderer;
