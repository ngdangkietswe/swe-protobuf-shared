package dev.ngdangkietswe.sweprotobufshared.proto.common;

import com.google.protobuf.GeneratedMessageV3;

import java.util.List;

/**
 * @author ngdangkietswe
 * @since 12/14/2024
 */

public abstract class GrpcMapper<F, T extends GeneratedMessageV3, B extends GeneratedMessageV3.Builder<B>> {

    public abstract B toGrpcBuilder(F from);

    public abstract T toGrpcMessage(F from);

    public List<T> toGrpcMessages(List<F> from) {
        return from.stream()
                .map(this::toGrpcMessage)
                .toList();
    }
}
