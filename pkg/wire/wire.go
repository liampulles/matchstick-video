package wire

import (
	goConfig "github.com/liampulles/go-config"

	"github.com/liampulles/matchstick-video/pkg/adapter/config"
	"github.com/liampulles/matchstick-video/pkg/adapter/db/sql"
	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	"github.com/liampulles/matchstick-video/pkg/adapter/http/json"
	"github.com/liampulles/matchstick-video/pkg/adapter/http/json/dto"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/driver/db"
	"github.com/liampulles/matchstick-video/pkg/driver/http/mux"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// CreateServer injects all the dependencies needed to create
// http.ServerFactory
func CreateServer(source goConfig.Source) (http.ServerFactory, error) {
	configStore, err := config.NewStoreImpl(
		source,
	)
	if err != nil {
		return nil, err
	}

	// --- NEXT TAP ---
	helperService := sql.NewHelperServiceImpl()
	databaseService, err := db.NewDatabaseServiceImpl(
		configStore,
	)
	if err != nil {
		return nil, err
	}
	inventoryItemConstructor := entity.NewInventoryItemConstructorImpl()
	muxWrapper := mux.NewWrapperImpl()

	// --- NEXT TAP ---
	inventoryRepository := sql.NewInventoryRepositoryImpl(
		databaseService,
		helperService,
		inventoryItemConstructor,
	)
	entityFactory := inventory.NewEntityFactoryImpl(
		inventoryItemConstructor,
	)
	entityModifier := inventory.NewEntityModifierImpl()
	ioMapper := mux.NewIOMapperImpl(
		muxWrapper,
	)

	// --- NEXT TAP ---
	inventoryService := inventory.NewServiceImpl(
		inventoryRepository,
		entityFactory,
		entityModifier,
	)
	decoderService := json.NewDecoderServiceImpl()
	encoderService := json.NewEncoderServiceImpl()
	responseFactory := http.NewResponseFactoryImpl()
	parameterConverter := http.NewParameterConverterImpl()
	dtoFactory := dto.NewFactoryImpl()
	if err != nil {
		return nil, err
	}
	handlerMapper := mux.NewHandlerMapperImpl(
		ioMapper,
	)

	// --- NEXT TAP ---
	inventoryController := http.NewInventoryControllerImpl(
		inventoryService,
		decoderService,
		encoderService,
		responseFactory,
		parameterConverter,
		dtoFactory,
	)
	serverConfiguration := mux.NewServerConfigurationImpl(
		configStore,
		handlerMapper,
		muxWrapper,
	)

	// --- NEXT TAP ---
	return http.NewServerFactoryImpl(
		inventoryController,
		serverConfiguration,
	), nil
}
