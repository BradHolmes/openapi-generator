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
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.time.LocalDate;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import org.openapitools.jackson.nullable.JsonNullable;
import com.fasterxml.jackson.annotation.JsonIgnore;
import org.openapitools.jackson.nullable.JsonNullable;
import java.util.NoSuchElementException;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;

/**
 * NullableClass
 */
@JsonPropertyOrder({
  NullableClass.JSON_PROPERTY_INTEGER_PROP,
  NullableClass.JSON_PROPERTY_NUMBER_PROP,
  NullableClass.JSON_PROPERTY_BOOLEAN_PROP,
  NullableClass.JSON_PROPERTY_STRING_PROP,
  NullableClass.JSON_PROPERTY_DATE_PROP,
  NullableClass.JSON_PROPERTY_DATETIME_PROP,
  NullableClass.JSON_PROPERTY_ARRAY_NULLABLE_PROP,
  NullableClass.JSON_PROPERTY_ARRAY_AND_ITEMS_NULLABLE_PROP,
  NullableClass.JSON_PROPERTY_ARRAY_ITEMS_NULLABLE,
  NullableClass.JSON_PROPERTY_OBJECT_NULLABLE_PROP,
  NullableClass.JSON_PROPERTY_OBJECT_AND_ITEMS_NULLABLE_PROP,
  NullableClass.JSON_PROPERTY_OBJECT_ITEMS_NULLABLE
})
@JsonTypeName("NullableClass")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class NullableClass extends HashMap<String, Object> {
  public static final String JSON_PROPERTY_INTEGER_PROP = "integer_prop";
  private JsonNullable<Integer> integerProp = JsonNullable.<Integer>undefined();

  public static final String JSON_PROPERTY_NUMBER_PROP = "number_prop";
  private JsonNullable<BigDecimal> numberProp = JsonNullable.<BigDecimal>undefined();

  public static final String JSON_PROPERTY_BOOLEAN_PROP = "boolean_prop";
  private JsonNullable<Boolean> booleanProp = JsonNullable.<Boolean>undefined();

  public static final String JSON_PROPERTY_STRING_PROP = "string_prop";
  private JsonNullable<String> stringProp = JsonNullable.<String>undefined();

  public static final String JSON_PROPERTY_DATE_PROP = "date_prop";
  private JsonNullable<LocalDate> dateProp = JsonNullable.<LocalDate>undefined();

  public static final String JSON_PROPERTY_DATETIME_PROP = "datetime_prop";
  private JsonNullable<OffsetDateTime> datetimeProp = JsonNullable.<OffsetDateTime>undefined();

  public static final String JSON_PROPERTY_ARRAY_NULLABLE_PROP = "array_nullable_prop";
  private JsonNullable<List<Object>> arrayNullableProp = JsonNullable.<List<Object>>undefined();

  public static final String JSON_PROPERTY_ARRAY_AND_ITEMS_NULLABLE_PROP = "array_and_items_nullable_prop";
  private JsonNullable<List<Object>> arrayAndItemsNullableProp = JsonNullable.<List<Object>>undefined();

  public static final String JSON_PROPERTY_ARRAY_ITEMS_NULLABLE = "array_items_nullable";
  private List<Object> arrayItemsNullable = null;

  public static final String JSON_PROPERTY_OBJECT_NULLABLE_PROP = "object_nullable_prop";
  private JsonNullable<Map<String, Object>> objectNullableProp = JsonNullable.<Map<String, Object>>undefined();

  public static final String JSON_PROPERTY_OBJECT_AND_ITEMS_NULLABLE_PROP = "object_and_items_nullable_prop";
  private JsonNullable<Map<String, Object>> objectAndItemsNullableProp = JsonNullable.<Map<String, Object>>undefined();

  public static final String JSON_PROPERTY_OBJECT_ITEMS_NULLABLE = "object_items_nullable";
  private Map<String, Object> objectItemsNullable = null;


  public NullableClass integerProp(Integer integerProp) {
    this.integerProp = JsonNullable.<Integer>of(integerProp);
    
    return this;
  }

   /**
   * Get integerProp
   * @return integerProp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonIgnore

  public Integer getIntegerProp() {
        return integerProp.orElse(null);
  }

  @JsonProperty(JSON_PROPERTY_INTEGER_PROP)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public JsonNullable<Integer> getIntegerProp_JsonNullable() {
    return integerProp;
  }
  
  @JsonProperty(JSON_PROPERTY_INTEGER_PROP)
  public void setIntegerProp_JsonNullable(JsonNullable<Integer> integerProp) {
    this.integerProp = integerProp;
  }

  public void setIntegerProp(Integer integerProp) {
    this.integerProp = JsonNullable.<Integer>of(integerProp);
  }


  public NullableClass numberProp(BigDecimal numberProp) {
    this.numberProp = JsonNullable.<BigDecimal>of(numberProp);
    
    return this;
  }

   /**
   * Get numberProp
   * @return numberProp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonIgnore

  public BigDecimal getNumberProp() {
        return numberProp.orElse(null);
  }

  @JsonProperty(JSON_PROPERTY_NUMBER_PROP)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public JsonNullable<BigDecimal> getNumberProp_JsonNullable() {
    return numberProp;
  }
  
  @JsonProperty(JSON_PROPERTY_NUMBER_PROP)
  public void setNumberProp_JsonNullable(JsonNullable<BigDecimal> numberProp) {
    this.numberProp = numberProp;
  }

  public void setNumberProp(BigDecimal numberProp) {
    this.numberProp = JsonNullable.<BigDecimal>of(numberProp);
  }


  public NullableClass booleanProp(Boolean booleanProp) {
    this.booleanProp = JsonNullable.<Boolean>of(booleanProp);
    
    return this;
  }

   /**
   * Get booleanProp
   * @return booleanProp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonIgnore

  public Boolean getBooleanProp() {
        return booleanProp.orElse(null);
  }

  @JsonProperty(JSON_PROPERTY_BOOLEAN_PROP)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public JsonNullable<Boolean> getBooleanProp_JsonNullable() {
    return booleanProp;
  }
  
  @JsonProperty(JSON_PROPERTY_BOOLEAN_PROP)
  public void setBooleanProp_JsonNullable(JsonNullable<Boolean> booleanProp) {
    this.booleanProp = booleanProp;
  }

  public void setBooleanProp(Boolean booleanProp) {
    this.booleanProp = JsonNullable.<Boolean>of(booleanProp);
  }


  public NullableClass stringProp(String stringProp) {
    this.stringProp = JsonNullable.<String>of(stringProp);
    
    return this;
  }

   /**
   * Get stringProp
   * @return stringProp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonIgnore

  public String getStringProp() {
        return stringProp.orElse(null);
  }

  @JsonProperty(JSON_PROPERTY_STRING_PROP)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public JsonNullable<String> getStringProp_JsonNullable() {
    return stringProp;
  }
  
  @JsonProperty(JSON_PROPERTY_STRING_PROP)
  public void setStringProp_JsonNullable(JsonNullable<String> stringProp) {
    this.stringProp = stringProp;
  }

  public void setStringProp(String stringProp) {
    this.stringProp = JsonNullable.<String>of(stringProp);
  }


  public NullableClass dateProp(LocalDate dateProp) {
    this.dateProp = JsonNullable.<LocalDate>of(dateProp);
    
    return this;
  }

   /**
   * Get dateProp
   * @return dateProp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonIgnore

  public LocalDate getDateProp() {
        return dateProp.orElse(null);
  }

  @JsonProperty(JSON_PROPERTY_DATE_PROP)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public JsonNullable<LocalDate> getDateProp_JsonNullable() {
    return dateProp;
  }
  
  @JsonProperty(JSON_PROPERTY_DATE_PROP)
  public void setDateProp_JsonNullable(JsonNullable<LocalDate> dateProp) {
    this.dateProp = dateProp;
  }

  public void setDateProp(LocalDate dateProp) {
    this.dateProp = JsonNullable.<LocalDate>of(dateProp);
  }


  public NullableClass datetimeProp(OffsetDateTime datetimeProp) {
    this.datetimeProp = JsonNullable.<OffsetDateTime>of(datetimeProp);
    
    return this;
  }

   /**
   * Get datetimeProp
   * @return datetimeProp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonIgnore

  public OffsetDateTime getDatetimeProp() {
        return datetimeProp.orElse(null);
  }

  @JsonProperty(JSON_PROPERTY_DATETIME_PROP)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public JsonNullable<OffsetDateTime> getDatetimeProp_JsonNullable() {
    return datetimeProp;
  }
  
  @JsonProperty(JSON_PROPERTY_DATETIME_PROP)
  public void setDatetimeProp_JsonNullable(JsonNullable<OffsetDateTime> datetimeProp) {
    this.datetimeProp = datetimeProp;
  }

  public void setDatetimeProp(OffsetDateTime datetimeProp) {
    this.datetimeProp = JsonNullable.<OffsetDateTime>of(datetimeProp);
  }


  public NullableClass arrayNullableProp(List<Object> arrayNullableProp) {
    this.arrayNullableProp = JsonNullable.<List<Object>>of(arrayNullableProp);
    
    return this;
  }

  public NullableClass addArrayNullablePropItem(Object arrayNullablePropItem) {
    if (this.arrayNullableProp == null || !this.arrayNullableProp.isPresent()) {
      this.arrayNullableProp = JsonNullable.<List<Object>>of(new ArrayList<>());
    }
    try {
      this.arrayNullableProp.get().add(arrayNullablePropItem);
    } catch (java.util.NoSuchElementException e) {
      // this can never happen, as we make sure above that the value is present
    }
    return this;
  }

   /**
   * Get arrayNullableProp
   * @return arrayNullableProp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonIgnore

  public List<Object> getArrayNullableProp() {
        return arrayNullableProp.orElse(null);
  }

  @JsonProperty(JSON_PROPERTY_ARRAY_NULLABLE_PROP)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public JsonNullable<List<Object>> getArrayNullableProp_JsonNullable() {
    return arrayNullableProp;
  }
  
  @JsonProperty(JSON_PROPERTY_ARRAY_NULLABLE_PROP)
  public void setArrayNullableProp_JsonNullable(JsonNullable<List<Object>> arrayNullableProp) {
    this.arrayNullableProp = arrayNullableProp;
  }

  public void setArrayNullableProp(List<Object> arrayNullableProp) {
    this.arrayNullableProp = JsonNullable.<List<Object>>of(arrayNullableProp);
  }


  public NullableClass arrayAndItemsNullableProp(List<Object> arrayAndItemsNullableProp) {
    this.arrayAndItemsNullableProp = JsonNullable.<List<Object>>of(arrayAndItemsNullableProp);
    
    return this;
  }

  public NullableClass addArrayAndItemsNullablePropItem(Object arrayAndItemsNullablePropItem) {
    if (this.arrayAndItemsNullableProp == null || !this.arrayAndItemsNullableProp.isPresent()) {
      this.arrayAndItemsNullableProp = JsonNullable.<List<Object>>of(new ArrayList<>());
    }
    try {
      this.arrayAndItemsNullableProp.get().add(arrayAndItemsNullablePropItem);
    } catch (java.util.NoSuchElementException e) {
      // this can never happen, as we make sure above that the value is present
    }
    return this;
  }

   /**
   * Get arrayAndItemsNullableProp
   * @return arrayAndItemsNullableProp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonIgnore

  public List<Object> getArrayAndItemsNullableProp() {
        return arrayAndItemsNullableProp.orElse(null);
  }

  @JsonProperty(JSON_PROPERTY_ARRAY_AND_ITEMS_NULLABLE_PROP)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public JsonNullable<List<Object>> getArrayAndItemsNullableProp_JsonNullable() {
    return arrayAndItemsNullableProp;
  }
  
  @JsonProperty(JSON_PROPERTY_ARRAY_AND_ITEMS_NULLABLE_PROP)
  public void setArrayAndItemsNullableProp_JsonNullable(JsonNullable<List<Object>> arrayAndItemsNullableProp) {
    this.arrayAndItemsNullableProp = arrayAndItemsNullableProp;
  }

  public void setArrayAndItemsNullableProp(List<Object> arrayAndItemsNullableProp) {
    this.arrayAndItemsNullableProp = JsonNullable.<List<Object>>of(arrayAndItemsNullableProp);
  }


  public NullableClass arrayItemsNullable(List<Object> arrayItemsNullable) {
    
    this.arrayItemsNullable = arrayItemsNullable;
    return this;
  }

  public NullableClass addArrayItemsNullableItem(Object arrayItemsNullableItem) {
    if (this.arrayItemsNullable == null) {
      this.arrayItemsNullable = new ArrayList<>();
    }
    this.arrayItemsNullable.add(arrayItemsNullableItem);
    return this;
  }

   /**
   * Get arrayItemsNullable
   * @return arrayItemsNullable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_ARRAY_ITEMS_NULLABLE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public List<Object> getArrayItemsNullable() {
    return arrayItemsNullable;
  }


  @JsonProperty(JSON_PROPERTY_ARRAY_ITEMS_NULLABLE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setArrayItemsNullable(List<Object> arrayItemsNullable) {
    this.arrayItemsNullable = arrayItemsNullable;
  }


  public NullableClass objectNullableProp(Map<String, Object> objectNullableProp) {
    this.objectNullableProp = JsonNullable.<Map<String, Object>>of(objectNullableProp);
    
    return this;
  }

  public NullableClass putObjectNullablePropItem(String key, Object objectNullablePropItem) {
    if (this.objectNullableProp == null || !this.objectNullableProp.isPresent()) {
      this.objectNullableProp = JsonNullable.<Map<String, Object>>of(new HashMap<>());
    }
    try {
      this.objectNullableProp.get().put(key, objectNullablePropItem);
    } catch (java.util.NoSuchElementException e) {
      // this can never happen, as we make sure above that the value is present
    }
    return this;
  }

   /**
   * Get objectNullableProp
   * @return objectNullableProp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonIgnore

  public Map<String, Object> getObjectNullableProp() {
        return objectNullableProp.orElse(null);
  }

  @JsonProperty(JSON_PROPERTY_OBJECT_NULLABLE_PROP)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public JsonNullable<Map<String, Object>> getObjectNullableProp_JsonNullable() {
    return objectNullableProp;
  }
  
  @JsonProperty(JSON_PROPERTY_OBJECT_NULLABLE_PROP)
  public void setObjectNullableProp_JsonNullable(JsonNullable<Map<String, Object>> objectNullableProp) {
    this.objectNullableProp = objectNullableProp;
  }

  public void setObjectNullableProp(Map<String, Object> objectNullableProp) {
    this.objectNullableProp = JsonNullable.<Map<String, Object>>of(objectNullableProp);
  }


  public NullableClass objectAndItemsNullableProp(Map<String, Object> objectAndItemsNullableProp) {
    this.objectAndItemsNullableProp = JsonNullable.<Map<String, Object>>of(objectAndItemsNullableProp);
    
    return this;
  }

  public NullableClass putObjectAndItemsNullablePropItem(String key, Object objectAndItemsNullablePropItem) {
    if (this.objectAndItemsNullableProp == null || !this.objectAndItemsNullableProp.isPresent()) {
      this.objectAndItemsNullableProp = JsonNullable.<Map<String, Object>>of(new HashMap<>());
    }
    try {
      this.objectAndItemsNullableProp.get().put(key, objectAndItemsNullablePropItem);
    } catch (java.util.NoSuchElementException e) {
      // this can never happen, as we make sure above that the value is present
    }
    return this;
  }

   /**
   * Get objectAndItemsNullableProp
   * @return objectAndItemsNullableProp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonIgnore

  public Map<String, Object> getObjectAndItemsNullableProp() {
        return objectAndItemsNullableProp.orElse(null);
  }

  @JsonProperty(JSON_PROPERTY_OBJECT_AND_ITEMS_NULLABLE_PROP)
  @JsonInclude(content = JsonInclude.Include.ALWAYS, value = JsonInclude.Include.USE_DEFAULTS)

  public JsonNullable<Map<String, Object>> getObjectAndItemsNullableProp_JsonNullable() {
    return objectAndItemsNullableProp;
  }
  
  @JsonProperty(JSON_PROPERTY_OBJECT_AND_ITEMS_NULLABLE_PROP)
  public void setObjectAndItemsNullableProp_JsonNullable(JsonNullable<Map<String, Object>> objectAndItemsNullableProp) {
    this.objectAndItemsNullableProp = objectAndItemsNullableProp;
  }

  public void setObjectAndItemsNullableProp(Map<String, Object> objectAndItemsNullableProp) {
    this.objectAndItemsNullableProp = JsonNullable.<Map<String, Object>>of(objectAndItemsNullableProp);
  }


  public NullableClass objectItemsNullable(Map<String, Object> objectItemsNullable) {
    
    this.objectItemsNullable = objectItemsNullable;
    return this;
  }

  public NullableClass putObjectItemsNullableItem(String key, Object objectItemsNullableItem) {
    if (this.objectItemsNullable == null) {
      this.objectItemsNullable = new HashMap<>();
    }
    this.objectItemsNullable.put(key, objectItemsNullableItem);
    return this;
  }

   /**
   * Get objectItemsNullable
   * @return objectItemsNullable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_OBJECT_ITEMS_NULLABLE)
  @JsonInclude(content = JsonInclude.Include.ALWAYS, value = JsonInclude.Include.USE_DEFAULTS)

  public Map<String, Object> getObjectItemsNullable() {
    return objectItemsNullable;
  }


  @JsonProperty(JSON_PROPERTY_OBJECT_ITEMS_NULLABLE)
  @JsonInclude(content = JsonInclude.Include.ALWAYS, value = JsonInclude.Include.USE_DEFAULTS)
  public void setObjectItemsNullable(Map<String, Object> objectItemsNullable) {
    this.objectItemsNullable = objectItemsNullable;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NullableClass nullableClass = (NullableClass) o;
    return equalsNullable(this.integerProp, nullableClass.integerProp) &&
        equalsNullable(this.numberProp, nullableClass.numberProp) &&
        equalsNullable(this.booleanProp, nullableClass.booleanProp) &&
        equalsNullable(this.stringProp, nullableClass.stringProp) &&
        equalsNullable(this.dateProp, nullableClass.dateProp) &&
        equalsNullable(this.datetimeProp, nullableClass.datetimeProp) &&
        equalsNullable(this.arrayNullableProp, nullableClass.arrayNullableProp) &&
        equalsNullable(this.arrayAndItemsNullableProp, nullableClass.arrayAndItemsNullableProp) &&
        Objects.equals(this.arrayItemsNullable, nullableClass.arrayItemsNullable) &&
        equalsNullable(this.objectNullableProp, nullableClass.objectNullableProp) &&
        equalsNullable(this.objectAndItemsNullableProp, nullableClass.objectAndItemsNullableProp) &&
        Objects.equals(this.objectItemsNullable, nullableClass.objectItemsNullable) &&
        super.equals(o);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(hashCodeNullable(integerProp), hashCodeNullable(numberProp), hashCodeNullable(booleanProp), hashCodeNullable(stringProp), hashCodeNullable(dateProp), hashCodeNullable(datetimeProp), hashCodeNullable(arrayNullableProp), hashCodeNullable(arrayAndItemsNullableProp), arrayItemsNullable, hashCodeNullable(objectNullableProp), hashCodeNullable(objectAndItemsNullableProp), objectItemsNullable, super.hashCode());
  }

  private static <T> int hashCodeNullable(JsonNullable<T> a) {
    if (a == null) {
      return 1;
    }
    return a.isPresent() ? Arrays.deepHashCode(new Object[]{a.get()}) : 31;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NullableClass {\n");
    sb.append("    ").append(toIndentedString(super.toString())).append("\n");
    sb.append("    integerProp: ").append(toIndentedString(integerProp)).append("\n");
    sb.append("    numberProp: ").append(toIndentedString(numberProp)).append("\n");
    sb.append("    booleanProp: ").append(toIndentedString(booleanProp)).append("\n");
    sb.append("    stringProp: ").append(toIndentedString(stringProp)).append("\n");
    sb.append("    dateProp: ").append(toIndentedString(dateProp)).append("\n");
    sb.append("    datetimeProp: ").append(toIndentedString(datetimeProp)).append("\n");
    sb.append("    arrayNullableProp: ").append(toIndentedString(arrayNullableProp)).append("\n");
    sb.append("    arrayAndItemsNullableProp: ").append(toIndentedString(arrayAndItemsNullableProp)).append("\n");
    sb.append("    arrayItemsNullable: ").append(toIndentedString(arrayItemsNullable)).append("\n");
    sb.append("    objectNullableProp: ").append(toIndentedString(objectNullableProp)).append("\n");
    sb.append("    objectAndItemsNullableProp: ").append(toIndentedString(objectAndItemsNullableProp)).append("\n");
    sb.append("    objectItemsNullable: ").append(toIndentedString(objectItemsNullable)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

