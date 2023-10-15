package utils

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"

	"project-management/shared"
)

func ReadExcelFile(ctx *fiber.Ctx, key string) ([][]string, error) {
	header, err := ctx.FormFile(key)

	if err != nil {
		shared.Logger.Error("ReadExcelFile", zap.Error(err))
		return nil, err
	}

	reader, err := header.Open()

	if err != nil {
		shared.Logger.Error("ReadExcelFile", zap.Error(err))
		return nil, err
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			shared.Logger.Warn("ReadExcelFile", zap.Error(err))
		}
	}(reader)

	excelFile, err := excelize.OpenReader(reader)

	if err != nil {
		shared.Logger.Error("ReadExcelFile", zap.Error(err))
		return nil, err
	}

	rows, err := excelFile.GetRows(excelFile.GetSheetName(0))

	if err != nil {
		shared.Logger.Error("ReadExcelFile", zap.Error(err))
		return nil, err
	}

	return rows, nil
}
