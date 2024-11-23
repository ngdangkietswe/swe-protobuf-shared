package dev.ngdangkietswe.sweprotobufshared.proto.domain;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.io.Serializable;
import java.util.List;
import java.util.UUID;

/**
 * @author ngdangkietswe
 * @since 11/23/2024
 */

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@Builder(setterPrefix = "set", builderMethodName = "newBuilder")
public class SweGrpcPrincipal implements Serializable {

    @JsonProperty("user_id")
    private UUID userId;
    private String username;
    private String email;
    private String token;
    private List<String> roles;
}
