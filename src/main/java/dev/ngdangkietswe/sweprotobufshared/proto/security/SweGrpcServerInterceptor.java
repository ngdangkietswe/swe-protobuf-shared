package dev.ngdangkietswe.sweprotobufshared.proto.security;

import com.fasterxml.jackson.databind.ObjectMapper;
import dev.ngdangkietswe.sweprotobufshared.proto.common.GrpcConstant;
import dev.ngdangkietswe.sweprotobufshared.proto.common.GrpcUtil;
import dev.ngdangkietswe.sweprotobufshared.proto.domain.SweGrpcPrincipal;
import io.grpc.Context;
import io.grpc.Contexts;
import io.grpc.Metadata;
import io.grpc.ServerCall;
import io.grpc.ServerCallHandler;
import io.grpc.ServerInterceptor;
import io.grpc.Status;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jws;
import io.jsonwebtoken.JwtParser;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.security.Keys;
import lombok.extern.log4j.Log4j2;
import org.apache.commons.lang3.StringUtils;

import java.util.Objects;

/**
 * @author ngdangkietswe
 * @since 11/23/2024
 * This class is used to intercept the incoming gRPC request and validate the JWT token
 */

@Log4j2
public class SweGrpcServerInterceptor implements ServerInterceptor {

    private static final ObjectMapper mapper = new ObjectMapper();

    private static final JwtParser jwtParser = Jwts.parserBuilder()
            .setSigningKey(Keys.hmacShaKeyFor(GrpcConstant.JWT_SECRET_KEY.getBytes()))
            .build();

    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(ServerCall<ReqT, RespT> serverCall,
                                                                 Metadata metadata,
                                                                 ServerCallHandler<ReqT, RespT> serverCallHandler) {
        Status status;
        var value = metadata.get(GrpcConstant.AUTHORIZATION_METADATA_KEY);
        if (StringUtils.isEmpty(value) || !value.startsWith(GrpcConstant.JWT_BEARER_PREFIX)) {
            status = Status.UNAUTHENTICATED.withDescription("Missing or invalid token");
        } else {
            try {
                var token = value.substring(GrpcConstant.JWT_BEARER_PREFIX.length()).trim();
                Jws<Claims> jws = jwtParser.parseClaimsJws(token);

                var principal = GrpcUtil.asGrpcPrincipal(jws);
                if (Objects.isNull(principal)) {
                    return Contexts.interceptCall(Context.current(), serverCall, metadata, serverCallHandler);
                }
                principal.setToken(token);

                // Get/Set user permission from metadata key to principal
                var sUserPermission = metadata.get(GrpcConstant.GRPC_USER_PERMISSION_METADATA_KEY);
                var userPermission = mapper.readValue(sUserPermission, SweGrpcPrincipal.UserPermission.class);
                principal.setUserPermission(userPermission);

                // Add principal to context, so we can access it later
                var context = Context.current().withValue(GrpcConstant.PRINCIPAL_CONTEXT_KEY, principal);
                return Contexts.interceptCall(context, serverCall, metadata, serverCallHandler);
            } catch (Exception e) {
                log.error("Error when parsing token: {}", e.getMessage());
                status = Status.UNAUTHENTICATED.withDescription(e.getMessage());
            }
        }

        serverCall.close(status, new Metadata());
        return new ServerCall.Listener<>() {
            // Do nothing
        };
    }
}
