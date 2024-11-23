package dev.ngdangkietswe.sweprotobufshared.proto.common;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.rpc.Code;
import dev.ngdangkietswe.sweprotobufshared.common.protobuf.Error;
import dev.ngdangkietswe.sweprotobufshared.common.protobuf.PageMetaData;
import dev.ngdangkietswe.sweprotobufshared.common.protobuf.Pageable;
import dev.ngdangkietswe.sweprotobufshared.proto.domain.SweGrpcPrincipal;
import dev.ngdangkietswe.sweprotobufshared.proto.exception.GrpcBadRequestException;
import dev.ngdangkietswe.sweprotobufshared.proto.exception.GrpcNotFoundException;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jws;
import org.apache.commons.lang3.StringUtils;
import org.springframework.lang.NonNull;

/**
 * @author ngdangkietswe
 * @since 11/23/2024
 */

public class GrpcUtil {

    private static final ObjectMapper mapper = new ObjectMapper();

    public static Error asError(Exception exception) {
        Error.Builder builder = Error.newBuilder().setMessage(getExceptionMessage(exception));

        if (exception instanceof GrpcNotFoundException) {
            builder
                    .setStatus(Code.NOT_FOUND.name())
                    .setCode(Code.NOT_FOUND_VALUE);
        } else if (exception instanceof GrpcBadRequestException ex) {
            builder
                    .setStatus(Code.INVALID_ARGUMENT.name())
                    .setCode(Code.INVALID_ARGUMENT_VALUE)
                    .putAllDetails(ex.getErrorDetails());
        } else {
            builder
                    .setStatus(Code.UNKNOWN.name())
                    .setCode(Code.UNKNOWN_VALUE);
        }

        return builder.build();
    }

    public static String getExceptionMessage(Exception exception) {
        return StringUtils.isNotEmpty(exception.getMessage())
                ? exception.getMessage()
                : "Unknown error";
    }

    public static PageMetaData asPageMetaData(Pageable pageable, long totalItems) {
        PageMetaData.Builder builder = PageMetaData.newBuilder()
                .setTotalElements(totalItems);

        if (pageable.getUnPaged()) {
            builder.setPage(1)
                    .setTotalPages(1)
                    .setSize((int) totalItems)
                    .setNext(false)
                    .setPrevious(false);
        } else {
            long totalPages = calculateTotalPages(totalItems, pageable.getSize());
            builder.setPage(pageable.getPage() + 1)
                    .setTotalPages(totalPages)
                    .setSize(pageable.getSize())
                    .setPrevious(pageable.getPage() > 0)
                    .setNext(pageable.getPage() + 1 < totalPages);
        }

        return builder.build();
    }

    public static long calculateOffset(int page, int size) {
        return (long) page * size;
    }

    public static long calculateTotalPages(long totalItems, int size) {
        return totalItems % size == 0 ? totalItems / size : totalItems / size + 1;
    }

    public static SweGrpcPrincipal asGrpcPrincipal(Jws<Claims> jws) {
        var body = jws.getBody();
        return mapper.convertValue(body.get("user"), SweGrpcPrincipal.class);
    }

    @NonNull
    public static SweGrpcPrincipal getGrpcPrincipal() {
        return GrpcConstant.PRINCIPAL_CONTEXT_KEY.get();
    }
}
