package dev.ngdangkietswe.sweprotobufshared.proto.exception;

/**
 * @author ngdangkietswe
 * @since 12/29/2024
 */

public class GrpcAccessDeniedException extends RuntimeException {

    public GrpcAccessDeniedException(String message) {
        super(message);
    }

    public GrpcAccessDeniedException() {
        this("Access denied");
    }
}
