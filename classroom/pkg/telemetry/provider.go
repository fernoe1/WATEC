package telemetry

import (
	"context"
	"time"

	"github.com/fernoe1/WATEC/classroom/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.40.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func fatal(err error) (
	*sdktrace.TracerProvider,
	*sdkmetric.MeterProvider,
	*sdklog.LoggerProvider,
	error,
) {
	return nil, nil, nil, err
}

func getTp(
	ctx context.Context,
	conn *grpc.ClientConn,
	res *resource.Resource,
) (*sdktrace.TracerProvider, error) {
	te, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, err
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(te),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	), nil
}

func getMp(
	ctx context.Context,
	conn *grpc.ClientConn,
	res *resource.Resource,
) (*sdkmetric.MeterProvider, error) {
	me, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, err
	}

	return sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(me,
			sdkmetric.WithInterval(15*time.Second),
		)),
		sdkmetric.WithResource(res),
	), nil
}

func getLp(
	ctx context.Context,
	conn *grpc.ClientConn,
	res *resource.Resource,
) (*sdklog.LoggerProvider, error) {
	le, err := otlploggrpc.New(ctx, otlploggrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, err
	}

	return sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(le)),
		sdklog.WithResource(res),
	), nil
}

func Init(ctx context.Context, cfg *config.Config) (
	*sdktrace.TracerProvider,
	*sdkmetric.MeterProvider,
	*sdklog.LoggerProvider,
	error,
) {
	res, err := resource.New(
		ctx,
		resource.WithAttributes(semconv.ServiceName(cfg.Telemetry.Name)),
	)
	if err != nil {
		return fatal(err)
	}

	conn, err := grpc.NewClient(
		cfg.Telemetry.Endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fatal(err)
	}

	tp, err := getTp(ctx, conn, res)
	if err != nil {
		return fatal(err)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	mp, err := getMp(ctx, conn, res)
	if err != nil {
		return fatal(err)
	}
	otel.SetMeterProvider(mp)

	lp, err := getLp(ctx, conn, res)
	if err != nil {
		return fatal(err)
	}
	global.SetLoggerProvider(lp)

	return tp, mp, lp, nil
}
