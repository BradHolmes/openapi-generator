/*
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import java.util.Map;
import java.util.HashMap;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonSubTypes;
import com.fasterxml.jackson.annotation.JsonTypeInfo;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.client.model.Quadrilateral;
import org.openapitools.client.model.Triangle;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;

import com.fasterxml.jackson.core.type.TypeReference;

import java.io.IOException;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashSet;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.JsonToken;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.MapperFeature;
import com.fasterxml.jackson.databind.SerializerProvider;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.deser.std.StdDeserializer;
import com.fasterxml.jackson.databind.ser.std.StdSerializer;
import org.openapitools.client.JSON;

@javax.annotation.processing.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
@JsonDeserialize(using = ShapeOrNull.ShapeOrNullDeserializer.class)
@JsonSerialize(using = ShapeOrNull.ShapeOrNullSerializer.class)
public class ShapeOrNull extends AbstractOpenApiSchema {
    private static final Logger log = Logger.getLogger(ShapeOrNull.class.getName());

    public static class ShapeOrNullSerializer extends StdSerializer<ShapeOrNull> {
        public ShapeOrNullSerializer(Class<ShapeOrNull> t) {
            super(t);
        }

        public ShapeOrNullSerializer() {
            this(null);
        }

        @Override
        public void serialize(ShapeOrNull value, JsonGenerator jgen, SerializerProvider provider) throws IOException, JsonProcessingException {
            jgen.writeObject(value.getActualInstance());
        }
    }

    public static class ShapeOrNullDeserializer extends StdDeserializer<ShapeOrNull> {
        public ShapeOrNullDeserializer() {
            this(ShapeOrNull.class);
        }

        public ShapeOrNullDeserializer(Class<?> vc) {
            super(vc);
        }

        @Override
        public ShapeOrNull deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
            JsonNode tree = jp.readValueAsTree();
            Object deserialized = null;
            ShapeOrNull newShapeOrNull = new ShapeOrNull();
            Map<String,Object> result2 = tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Map<String, Object>>() {});
            String discriminatorValue = (String)result2.get("shapeType");
            switch (discriminatorValue) {
                case "Quadrilateral":
                    deserialized = tree.traverse(jp.getCodec()).readValueAs(Quadrilateral.class);
                    newShapeOrNull.setActualInstance(deserialized);
                    return newShapeOrNull;
                case "Triangle":
                    deserialized = tree.traverse(jp.getCodec()).readValueAs(Triangle.class);
                    newShapeOrNull.setActualInstance(deserialized);
                    return newShapeOrNull;
                default:
                    log.log(Level.WARNING, String.format("Failed to lookup discriminator value `%s` for ShapeOrNull. Possible values: Quadrilateral Triangle", discriminatorValue));
            }

            boolean typeCoercion = ctxt.isEnabled(MapperFeature.ALLOW_COERCION_OF_SCALARS);
            int match = 0;
            JsonToken token = tree.traverse(jp.getCodec()).nextToken();
            // deserialize Quadrilateral
            try {
                boolean attemptParsing = true;
                // ensure that we respect type coercion as set on the client ObjectMapper
                if (Quadrilateral.class.equals(Integer.class) || Quadrilateral.class.equals(Long.class) || Quadrilateral.class.equals(Float.class) || Quadrilateral.class.equals(Double.class) || Quadrilateral.class.equals(Boolean.class) || Quadrilateral.class.equals(String.class)) {
                    attemptParsing = typeCoercion;
                    if (!attemptParsing) {
                        attemptParsing |= ((Quadrilateral.class.equals(Integer.class) || Quadrilateral.class.equals(Long.class)) && token == JsonToken.VALUE_NUMBER_INT);
                        attemptParsing |= ((Quadrilateral.class.equals(Float.class) || Quadrilateral.class.equals(Double.class)) && token == JsonToken.VALUE_NUMBER_FLOAT);
                        attemptParsing |= (Quadrilateral.class.equals(Boolean.class) && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE));
                        attemptParsing |= (Quadrilateral.class.equals(String.class) && token == JsonToken.VALUE_STRING);
                        attemptParsing |= (token == JsonToken.VALUE_NULL);
                    }
                }
                if (attemptParsing) {
                    deserialized = tree.traverse(jp.getCodec()).readValueAs(Quadrilateral.class);
                    // TODO: there is no validation against JSON schema constraints
                    // (min, max, enum, pattern...), this does not perform a strict JSON
                    // validation, which means the 'match' count may be higher than it should be.
                    match++;
                    log.log(Level.FINER, "Input data matches schema 'Quadrilateral'");
                }
            } catch (Exception e) {
                // deserialization failed, continue
                log.log(Level.FINER, "Input data does not match schema 'Quadrilateral'", e);
            }

            // deserialize Triangle
            try {
                boolean attemptParsing = true;
                // ensure that we respect type coercion as set on the client ObjectMapper
                if (Triangle.class.equals(Integer.class) || Triangle.class.equals(Long.class) || Triangle.class.equals(Float.class) || Triangle.class.equals(Double.class) || Triangle.class.equals(Boolean.class) || Triangle.class.equals(String.class)) {
                    attemptParsing = typeCoercion;
                    if (!attemptParsing) {
                        attemptParsing |= ((Triangle.class.equals(Integer.class) || Triangle.class.equals(Long.class)) && token == JsonToken.VALUE_NUMBER_INT);
                        attemptParsing |= ((Triangle.class.equals(Float.class) || Triangle.class.equals(Double.class)) && token == JsonToken.VALUE_NUMBER_FLOAT);
                        attemptParsing |= (Triangle.class.equals(Boolean.class) && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE));
                        attemptParsing |= (Triangle.class.equals(String.class) && token == JsonToken.VALUE_STRING);
                        attemptParsing |= (token == JsonToken.VALUE_NULL);
                    }
                }
                if (attemptParsing) {
                    deserialized = tree.traverse(jp.getCodec()).readValueAs(Triangle.class);
                    // TODO: there is no validation against JSON schema constraints
                    // (min, max, enum, pattern...), this does not perform a strict JSON
                    // validation, which means the 'match' count may be higher than it should be.
                    match++;
                    log.log(Level.FINER, "Input data matches schema 'Triangle'");
                }
            } catch (Exception e) {
                // deserialization failed, continue
                log.log(Level.FINER, "Input data does not match schema 'Triangle'", e);
            }

            if (match == 1) {
                ShapeOrNull ret = new ShapeOrNull();
                ret.setActualInstance(deserialized);
                return ret;
            }
            throw new IOException(String.format("Failed deserialization for ShapeOrNull: %d classes match result, expected 1", match));
        }

        /**
         * Handle deserialization of the 'null' value.
         */
        @Override
        public ShapeOrNull getNullValue(DeserializationContext ctxt) throws JsonMappingException {
            return null;
        }
    }

    // store a list of schema names defined in oneOf
    public static final Map<String, Class<?>> schemas = new HashMap<>();

    public ShapeOrNull() {
        super("oneOf", Boolean.TRUE);
    }

    public ShapeOrNull(Quadrilateral o) {
        super("oneOf", Boolean.TRUE);
        setActualInstance(o);
    }

    public ShapeOrNull(Triangle o) {
        super("oneOf", Boolean.TRUE);
        setActualInstance(o);
    }

    static {
        schemas.put("Quadrilateral", Quadrilateral.class);
        schemas.put("Triangle", Triangle.class);
        JSON.registerDescendants(ShapeOrNull.class, Collections.unmodifiableMap(schemas));
        // Initialize and register the discriminator mappings.
        Map<String, Class<?>> mappings = new HashMap<String, Class<?>>();
        mappings.put("Quadrilateral", Quadrilateral.class);
        mappings.put("Triangle", Triangle.class);
        mappings.put("ShapeOrNull", ShapeOrNull.class);
        JSON.registerDiscriminator(ShapeOrNull.class, "shapeType", mappings);
    }

    @Override
    public Map<String, Class<?>> getSchemas() {
        return ShapeOrNull.schemas;
    }

    /**
     * Set the instance that matches the oneOf child schema, check
     * the instance parameter is valid against the oneOf child schemas:
     * Quadrilateral, Triangle
     *
     * It could be an instance of the 'oneOf' schemas.
     * The oneOf child schemas may themselves be a composed schema (allOf, anyOf, oneOf).
     */
    @Override
    public void setActualInstance(Object instance) {
        if (instance == null) {
           super.setActualInstance(instance);
           return;
        }

        if (JSON.isInstanceOf(Quadrilateral.class, instance, new HashSet<Class<?>>())) {
            super.setActualInstance(instance);
            return;
        }

        if (JSON.isInstanceOf(Triangle.class, instance, new HashSet<Class<?>>())) {
            super.setActualInstance(instance);
            return;
        }

        throw new RuntimeException("Invalid instance type. Must be Quadrilateral, Triangle");
    }

    /**
     * Get the actual instance, which can be the following:
     * Quadrilateral, Triangle
     *
     * @return The actual instance (Quadrilateral, Triangle)
     */
    @Override
    public Object getActualInstance() {
        return super.getActualInstance();
    }

    /**
     * Get the actual instance of `Quadrilateral`. If the actual instanct is not `Quadrilateral`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `Quadrilateral`
     * @throws ClassCastException if the instance is not `Quadrilateral`
     */
    public Quadrilateral getQuadrilateral() throws ClassCastException {
        return (Quadrilateral)super.getActualInstance();
    }

    /**
     * Get the actual instance of `Triangle`. If the actual instanct is not `Triangle`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `Triangle`
     * @throws ClassCastException if the instance is not `Triangle`
     */
    public Triangle getTriangle() throws ClassCastException {
        return (Triangle)super.getActualInstance();
    }

}

