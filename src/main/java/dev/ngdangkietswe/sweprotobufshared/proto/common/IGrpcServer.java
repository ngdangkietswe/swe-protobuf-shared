package dev.ngdangkietswe.sweprotobufshared.proto.common;

import dev.ngdangkietswe.sweprotobufshared.proto.domain.SweGrpcPrincipal;
import io.grpc.stub.StreamObserver;

import java.util.function.BiFunction;
import java.util.function.Function;

/**
 * @author ngdangkietswe
 * @since 11/21/2024
 */

public interface IGrpcServer {

    static <I, O> void execute(
            I req,
            StreamObserver<O> streamObserver,
            Function<I, O> success,
            Function<Exception, O> error
    ) {
        try {
            streamObserver.onNext(success.apply(req));
            streamObserver.onCompleted();
        } catch (Exception e) {
            streamObserver.onNext(error.apply(e));
            streamObserver.onCompleted();
            throw e;
        }
    }

    static <I, O> void execute(
            I req,
            StreamObserver<O> streamObserver,
            BiFunction<I, SweGrpcPrincipal, O> success,
            Function<Exception, O> error
    ) {
        try {
            var principal = GrpcUtil.getGrpcPrincipal();
            streamObserver.onNext(success.apply(req, principal));
            streamObserver.onCompleted();
        } catch (Exception e) {
            streamObserver.onNext(error.apply(e));
            streamObserver.onCompleted();
            throw e;
        }
    }
}
