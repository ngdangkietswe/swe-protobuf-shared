package dev.ngdangkietswe.sweprotobufshared.proto.common;

import com.google.common.net.HttpHeaders;
import dev.ngdangkietswe.sweprotobufshared.proto.domain.SweGrpcPrincipal;
import io.grpc.Context;
import io.grpc.Metadata;

import static io.grpc.Metadata.ASCII_STRING_MARSHALLER;

/**
 * @author ngdangkietswe
 * @since 11/23/2024
 */

public class GrpcConstant {

    public static final String JWT_BEARER_PREFIX = "Bearer";
    public static final String JWT_SECRET_KEY = "jwts3cr3tjwts3cr3tjwts3cr3tjwts3cr3tjwts3cr3tjwts3cr3tjwts3cr3t";

    public static final Metadata.Key<String> AUTHORIZATION_METADATA_KEY = Metadata.Key
            .of(HttpHeaders.AUTHORIZATION, ASCII_STRING_MARSHALLER);
    public static final Context.Key<SweGrpcPrincipal> PRINCIPAL_CONTEXT_KEY = Context.key(HttpHeaders.AUTHORIZATION);
}
