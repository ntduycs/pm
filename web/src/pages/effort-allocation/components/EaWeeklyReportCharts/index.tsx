import { StyledEaWeeklyReportCharts } from '@pm/pages/effort-allocation/components/EaWeeklyReportCharts/styles.ts';
import { ArcElement, Chart, ChartData, ChartOptions, Legend, Title, Tooltip } from 'chart.js';
import { effortAllocationState } from '@pm/atoms';
import { useRecoilValue } from 'recoil';
import { Pie } from 'react-chartjs-2';
import { EffortAllocation } from '@pm/common/constants';
import { Col, Row } from 'antd';

Chart.register(ArcElement, Tooltip, Legend, Title);
Chart.overrides.pie.plugins.legend = {
  ...Chart.overrides.pie.plugins.legend,
  display: false,
};

export const EaWeeklyReportCharts = () => {
  const eaState = useRecoilValue(effortAllocationState);

  const pieLabelCategoryData: ChartData<'pie', number[], unknown> = {
    labels: Object.values(EffortAllocation.Category).map((item) => item.label),
    datasets: [
      {
        data: Object.values(EffortAllocation.Category).map((item) => {
          return eaState.reduce((total, member) => {
            const effort = member.efforts.find((effort) => effort.category === item.label);
            return total + (effort?.time ?? 0);
          }, 0);
        }),
        backgroundColor: Object.values(EffortAllocation.Category).map(
          (item) => item.backgroundColor,
        ),
        borderColor: Object.values(EffortAllocation.Category).map((item) => item.borderColor),
      },
    ],
  };

  const pieProductivityData: ChartData<'pie', number[], unknown> = {
    labels: ['Production Time', 'Non-production Time'],
    datasets: [
      {
        data: [
          eaState.reduce((total, member) => {
            return (
              total +
              member.efforts.reduce((totalEffort, effort) => {
                if (effort.is_product_time) {
                  return totalEffort + effort.time;
                }
                return totalEffort;
              }, 0)
            );
          }, 0),
          eaState.reduce((total, member) => {
            return (
              total +
              member.efforts.reduce((totalEffort, effort) => {
                if (!effort.is_product_time) {
                  return totalEffort + effort.time;
                }
                return totalEffort;
              }, 0)
            );
          }, 0),
        ],
        backgroundColor: ['#b7eb8f', '#ffa39e'],
        borderColor: ['#52c41a', '#f5222d'],
      },
    ],
  };

  const pieLabelCategoryOptions: ChartOptions<'pie'> = {
    responsive: true,
    normalized: true,
    plugins: {
      title: {
        display: true,
        text: 'Sum per Category (h)',
      },
    },
  };

  const pieProductivityOptions: ChartOptions<'pie'> = {
    responsive: true,
    normalized: true,
    plugins: {
      title: {
        display: true,
        text: 'Production vs. Non-production Time (h)',
      },
    },
  };

  return (
    eaState.length > 0 && (
      <StyledEaWeeklyReportCharts>
        <Row>
          <Col span={8}>
            <Pie
              data={pieLabelCategoryData}
              options={pieLabelCategoryOptions}
            />
          </Col>
          <Col span={8}>
            <Pie
              data={pieProductivityData}
              options={pieProductivityOptions}
            />
          </Col>
        </Row>
      </StyledEaWeeklyReportCharts>
    )
  );
};
