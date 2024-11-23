package dev.ngdangkietswe.sweprotobufshared.proto.exception;

/**
 * @author ngdangkietswe
 * @since 11/23/2024
 */

public class GrpcNotFoundException extends RuntimeException {

    public GrpcNotFoundException(String message) {
        super(message);
    }

    public GrpcNotFoundException(Class<?> clazz, String key, Object value) {
        super(String.format("Not found %s with %s = %s", clazz.getSimpleName(), key, value));
    }
}
