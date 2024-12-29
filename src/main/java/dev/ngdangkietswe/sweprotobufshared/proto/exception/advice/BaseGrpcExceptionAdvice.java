package dev.ngdangkietswe.sweprotobufshared.proto.exception.advice;

import dev.ngdangkietswe.sweprotobufshared.proto.exception.GrpcAccessDeniedException;
import io.grpc.Status;
import org.lognet.springboot.grpc.recovery.GRpcExceptionHandler;
import org.lognet.springboot.grpc.recovery.GRpcExceptionScope;

/**
 * @author ngdangkietswe
 * @since 12/29/2024
 */

public abstract class BaseGrpcExceptionAdvice {

    @GRpcExceptionHandler
    protected Status handleAccessDenied(GrpcAccessDeniedException ex, GRpcExceptionScope scope) {
        return Status.PERMISSION_DENIED.withDescription(ex.getMessage());
    }
}
