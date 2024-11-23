package dev.ngdangkietswe.sweprotobufshared.proto.exception;

import lombok.Getter;

import java.util.Map;

/**
 * @author ngdangkietswe
 * @since 11/23/2024
 */

@Getter
public class GrpcBadRequestException extends RuntimeException {

    private final Map<String, String> errorDetails;

    public GrpcBadRequestException(String message, Map<String, String> errorDetails) {
        super(message);
        this.errorDetails = errorDetails;
    }
}
