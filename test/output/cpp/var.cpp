#include <map>
#include <list>
#include <string>
#include <iostream>
#define NAMEOF(variable) ((void)variable, #variable)

using namespace std;

template<class T>
string foo(T var, const char* varname)
{
    // std::cout << varname << std::endl;
    return varname;
}

#define FOO(var) foo(var, NAMEOF(var))

// namespace var {

    // struct var { std::string type; void* data; };
    // std::unordered_map<std::type_index, std::string> type_names;

    // type_names[std::type_index(typeid(int))] = "int";
    // type_names[std::type_index(typeid(double))] = "double";

    enum varType { 
        pointerType, 
        intType, 
        boolType,
        charType,
        floatType,
        stringType,
        structType,
        objectType,
        arrayType
    };

    map<string, varType> typeidMap = {
        { "i", intType },
        { "PKc", stringType },
    };

    // static map<varType, function<void(void*)>> deleteMap = {
    //     { intType,
    //         [](void* lambdaptr) {
    //             printf("int decons\n");
    //             delete (int*)lambdaptr; 
    //         }
    //     }, 
    //     { stringType,
    //         [](void* lambdaptr) {
    //             printf("string decons; Value: %p\n", lambdaptr);
    //             delete (string*)lambdaptr; 
    //             printf("string decons; Value: %p\n", lambdaptr);
    //         }
    //     }
    // };
    // deleteMap[intType] = []() { printf("wtf"); };
    //  = {
    //     { intType,         
    //         [](void* lambdaData) { 
    //             printf("running this shit\n"); 
    //             delete (int*)lambdaData; 
    //         } 
    //     },
    //     { "PKc", stringType },
    // };



    // FIXME: this only supports homogenous type operations right now
    class var {
        // TODO: I'm sure we'll need to add more later on

        // simple gc bleh
        // map<string, 
        // map from pointer to var of all declared
        // map from pointer to var of in use
        // if that pointer is not in use anymore, i.e - redeclared, delete from both lists

        private:
            string name;
            varType type;
            // TODO: FIXME: change this to be a smart pointer using shared_ptr<void>()
            void* data;
        
        public:
            // ~var() {
            //     deleteMap[type](data);
            // }

            // ~var() {
            void deallocate() {
                switch(type) {
                    // FIXME: not sure how to handle this
                    // case pointerType:

                    case intType: {
                        cout << "int decons; Name: " << name << " Type: " << type << " Value: " << *(int*)data << " Pointer: " << data << endl;
                        // delete (int*)data;
                        break;
                    }
                        
                    case stringType: {
                        cout << "string decons; Name: " << name << " Type: " << type << " Value: " << *(string*)data << " Pointer: " << data << endl;
                        // delete (string*)data;
                        break;
                    }

                    case boolType: {
                        cout << "bool decons; Name: " << name << " Type: " << type << " Value: " << *(bool*)data << " Pointer: " << data << endl;
                        delete (bool*)data;
                        break;
                    }

                    case charType: {
                        cout << "char decons; Name: " << name << " Type: " << type << " Value: " << *(char*)data << " Pointer: " << data << endl;
                        delete (char*)data;
                        break;
                    }
                    
                    case floatType: {
                        cout << "float decons; Name: " << name << " Type: " << type << " Value: " << *(float*)data << " Pointer: " << data << endl;
                        delete (float*)data;
                        break;
                    }

                    case objectType: {
                        cout << "object decons; Name: " << name << " Type: " << type << " Value: " << *this << " Pointer: " << data << endl;
                        delete (map<string, var>*)data;
                        break;
                    }
                    
                    default:
                        printf("don't know how to deallocate; Type: %u Value: %p\n", type, data);
                }

                    // FIXME: finish when structs are implemented
                    // case arrayType:
                    // case structType:
                    // case objectType:
            }

            // FIXME: switch ALL constructors to use pass-by-reference
            // TODO: this need to assert a default type
            // var() : type(intType), data(new int(0)) {}
            var(void) : type(objectType), data(new map<string, var>) {}
            var(void* value) : type(pointerType), data(value) {}
            var(int value) {
                // name = FOO(value);
                // if (type != 0) {
                //     // var::~var();
                //     printf("dater %d\n", *(int*)data);
                //     printf("type %d\n", type);
                //     // delete data;
                //     printf("data was not null\n");
                // }
                type = intType;
                // data = new int(value);
                *(int*)data = value;
                cout << "int cons; Name: " << name << " Type: " << type << " Value: " << value << " Pointer: " << data << endl;
                //printf("int cons\n");
            }
            var(bool value) : type(boolType), data(new bool(value)) {
                cout << "bool cons; Name: " << name << " Type: " << type << " Value: " << value << " Pointer: " << data << endl;
            }
            var(char value) : type(charType), data(new char(value)) {}
            var(float value) : type(floatType), data(new float(value)) {
                cout << "float cons; Name: " << name << " Type: " << type << " Value: " << value << " Pointer: " << data << endl;
            }
            var(double value) : type(floatType), data(new float(value)) {
                cout << "float cons; Name: " << name << " Type: " << type << " Value: " << value << " Pointer: " << data << endl;
            }
            // all string literal constructions are going in here
            var(const char* value) : type(stringType), data(new string(value)) {
                cout << "string cons; Name: " << name << " Type: " << type << " Value: \"" << value << "\" Pointer: " << data << endl;
            }
            var(string value) : type(stringType), data(new string(value)) {
                cout << "string cons; Name: " << name << " Type: " << type << " Value: \"" << value << "\" Pointer: " << data << endl;
            }

            var(initializer_list<var> propList) : type(objectType) {

                    printf("i am here ?? \n");
                if (propList.size() % 2 != 0) {
                    cout << "ERROR: invalid amount of arguments to object" << endl;
                    exit(9);
                }

                map<string, var> object;

                int i = 0;
                var lastItem;
                for (auto prop : propList) {
                    if (i % 2 == 1) {
                        // TODO: could think about supporting more than string here but that will involve many more more checks
                        object[*(string*)lastItem.Value()] = prop;
                    } else {
                        lastItem = prop;
                    }

                    i++;
                }

                // something weird is happening here....
                printf("wtffff\n");
                data = new map<string, var>(object);
                // FIXME: ... somehow this will work ...
                // *(map<string, var>*)data = object;
            }

            // TODO: will have to do something special here, maybe code generation?
            // var(struct value) : type(structType), data(&value) {}
            // TODO: not sure if you can do this with a map, might have to copy everything over
            // var(map<string, var> value) : type(objectType), data(new map<string, var>(value)) {
            //     //printf("obj cons\n");
            // }
            // FIXME: might take this out, kind of unsafe
            var(varType iType, void* iData) : type(iType), data(iData) {
                //printf("void*\n");
            }

            varType Type(void) const {
                return type;
            }

            void* Value(void) const {
                //printf("value\n");
                return data;
            }

            var& operator[](string attribute) {
                printf("calling this\n");
                if (type == objectType) {
                    return (*(map<string, var>*)data)[attribute];
                }
                 else {
                    type = objectType;
                    map<string, var> object;
                    object[attribute] = 0;
                    
                    data = (void*)&object;
                    return (*(map<string, var>*)data)[attribute];
                }
            }

            void operator+=(const int right) {
                // printf("+= var int\n");
                *(int*)data += right;
            }

            int operator*(const var& right) {
                // printf("* var var\n");
                return *(int*)data * *(int*)right.data;
            }

            void operator=(const int right) {
                if (type == intType) {
                    *(int*)data = right;
                } else {
                    // var::~var();
                    deallocate();
                    printf("int cons; Type: %u Value: %p\n", type, data);
                    type = intType;
                    // data = new int(right);
                    *(int*)data = right;
                }
            }

            void operator=(const double right) {
                if (type == floatType) {
                    *(float*)data = right;
                } else {
                    // var::~var();
                    deallocate();
                    printf("float cons; Type: %u Value: %p\n", type, data);
                    type = floatType;
                    // data = new float(right);
                    *(float*)data = right;
                }
            }

            void operator=(const char* right) {
                if (type == stringType) {
                    *(string*)data = right;
                } else {
                    // var::~var();
                    deallocate();
                    cout << "string cons; Name: " << name << " Type: " << type << " Value: \"" << right << "\" Pointer: " << data << endl;
                    type = stringType;
                    // data = new string(right);
                    *(string*)data = right;
                }
            }

            void operator=(const bool right) {
                if (type == boolType) {
                    *(bool*)data = right;
                } else {
                    // var::~var();
                    deallocate();
                    printf("bool cons; Type: %u Value: %p\n", type, data);
                    type = boolType;
                    *(bool*)data = right;
                }
            }

            void operator=(initializer_list<var> propList) {
                deallocate();
                // cout << "object cons; Name: " << name << " Type: " << type << " Value: " << propList << " Pointer: " << data << endl;
                cout << "object cons; Name: " << name << " Type: " << type << " Pointer: " << data << endl;
                type = objectType;
                printf("i got here\n");
                var thing = propList;
                // cout << thing << endl;
                // data = var(propList).data;
                data = thing.data;
                // printf("i got here too\n");
            }

            // template <typename T>
            // void operator=(T right) {
            //     // cout << right << " type: " << typeid(right).name() << endl;
            //     // cout << typeidMap[typeid(right).name()] << endl;

            //     if (type == typeidMap[typeid(right).name()]) {
            //         data = new T(right);
            //         // data = (T*)right;
            //     } 
            //     else {
            //         var::~var();
            //         // printf("type: %d\n", type);
            //         // printf("type id: %d\n", typeidMap[typeid(right).name()]);
            //         // printf("type id: %s\n", typeid(right).name());
            //         type = typeidMap[typeid(right).name()];
            //         // printf("type again: %d\n", type);
            //         // cout << "right: " << right << endl;

            //         // data = new T(right);
            //         // data = new string("hi");
            //         // var datavalue = new var(right);
            //         data = (new var(right)) -> data;

            //         // if (type == stringType) {
            //             // var thing = right;
            //             // data = thing.Value;
            //         // }
            //     }

            //     // if (type != typeidMap[typeid(right).name()]) {
            //     //     var::~var();
            //     //     type = typeidMap[typeid(right).name()];
            //     // }
            //     // // data = new T(right);

            //     // switch(type) {
            //     //     case intType:
            //     //         data = new int(right);
            //     //     case stringType:
            //     //         data = new string(right);
            //     //     default:
            //     //         printf("couldn't");
            //     // }
            // }

            // TODO: FIXME: could not get this to work
            // var& operator=(initializer_list<var> propList) {
            //     //printf("im in");

            //     int propListLen = propList.size();
            //     if (propListLen % 2 != 0) {
            //         cout << "ERROR: invalid amount of arguments to object" << endl;
            //         exit(9);
            //     }

            //     map<string, var> object;

            //     cout << object << endl;

            //     // TODO: ok this is kinda fucking hacky but w/e
            //     // var thing[propListLen];
            //     int i = 0;
            //     var lastItem;
            //     for (auto prop : propList) {
            //         cout << prop << endl;

            //         if (i % 2 == 1) {
            //             object[*(string*)lastItem.Value()] = prop;
            //         } else {
            //             lastItem = prop;
            //         }

            //         i++;
            //     }

            //     cout << object << endl;

            //     // for (int i = 0; i < propListLen; i+=2) {
            //     //     // cout << "i " << i << thing[i] << endl;
            //     //     // cout << "i " << i+1 << thing[i+1] << endl;
            //     //     // FIXME: -SUPER- unsafe, need to make sure they are strings
            //     //     object[*(string*)thing[i].Value()] = thing[i+1];
            //     //     cout << object << endl;
            //     // }


            //     data(new map<string, var>(object));
            // }

            friend ostream& operator<<(ostream& stream, var& v) {
                switch(v.type) {
                    case intType:
                        // printf("printing int\n");
                        return stream << *(int*)v.data;

                    case boolType:
                        if (*(bool*)v.data) {
                            return stream << "true";
                        }
                        return stream << "false";

                    case charType:
                        return stream << *(char*)v.data;

                    case floatType:
                        return stream << *(float*)v.data;

                    case stringType:
                        // cout << "printing string" << endl;;
                        return stream << *(string*)v.data;

                    case objectType: {
                        int counter = 0;
                        map<string, var> objectMap = *(map<string, var>*)v.data;
                        stream << "{ ";
                        for (auto property : objectMap) {
                            // stream << property.first << " " << property.second.first << " " << property.second.second << "\n";
                            stream << property.first << ": " << property.second;

                            if (counter < objectMap.size()-1) {
                                stream << ", ";
                            }
                            counter++;
                        }
                        return stream << " }";
                    }

                    default:
                        printf("wtf to do Type: %u\n", v.type);
                }

                return stream;
            }
    };

    map<void*, var> declared;
    map<void*, var> inuse;

    int gc(void) {
        printf("gc time...\n");

        for (auto memloc : declared) {
            cout << memloc.first << " " << memloc.second << endl;
        }

        return 0;
    }

    // TODO: for right now, instead of doing the map[string]function to figure out the value
    // https://stackoverflow.com/questions/4972795/how-do-i-typecast-with-type-info
    // https://stackoverflow.com/questions/2136998/using-a-stl-map-of-function-pointer
    // ostream& operator<<(ostream& stream, const var& v){
    //     switch(v.Type()) {
    //         case intType:
    //             return stream << *(int*)v.Value();

    //         case boolType:
    //             if (*(bool*)v.Value()) {
    //                 return stream << "true";
    //             }
    //             return stream << "false";

    //         case charType:
    //             return stream << *(char*)v.Value();

    //         case floatType:
    //             return stream << *(float*)v.Value();

    //         case stringType:
    //             return stream << *(string*)v.Value();

    //         case objectType: {
    //             int counter = 0;
    //             map<string, var> objectMap = *(map<string, var>*)v.Value();
    //             int mapLen = objectMap.size();
    //             stream << "{ ";
    //             for (auto property : objectMap) {
    //                 // stream << property.first << " " << property.second.first << " " << property.second.second << "\n";
    //                 stream << property.first << ": " << property.second;

    //                 if (counter < mapLen-1) {
    //                     stream << ", ";
    //                 }
    //                 counter++;
    //             }
    //             return stream << " }";
    //         }

    //         default:
    //             //printf("wtf to do Type: %u\n", v.Type());
    //     }

    //     return stream;
    // }

    // Interger operations
    int operator+(const int left, const var& right) {
        // printf("+ int var\n");
        return left + *(int*)right.Value();
    }

    int operator+=(int left, const var& right) {
        // printf("+= int var\n");
        return left += *(int*)right.Value();
    }

    // int operator+=(const var& right, int left) {
    //     printf("+= var int\n");
    //     return left += *(int*)right.Value();
    // }

    int operator+=(const var& left, const var& right) {
        // printf("+= var var\n");
        return *(int*)left.Value() += *(int*)right.Value();
    }

    int operator*(const int left, const var& right) {
        // printf("* int var\n");
        return left * *(int*)right.Value();
    }

    bool operator+(const bool left, const var& right) {
        return left || *(bool*)right.Value();
    }

    // TODO: not sure about this one for now
    // char operator+(const char left, const var& right) {
    //     return left || *(bool*)right.Value();
    // }

    float operator+(const float left, const var& right) {
        return left + *(float*)right.Value();
    }

    float operator+(const double left, const var& right) {
        return left + *(float*)right.Value();
    }

    // String/Char* operations: convert char* to string with all of these functions
    string operator+(const char* left, const var& right) {
        return left + *(string*)right.Value();
    }

    var operator+(const var& left, const char* right) {
        return var{ *(string*)left.Value() + right };
    }

    // Generic constructor for right side value
    template <typename T>
    var operator+(const var& left, T right) {
        // FIXME: this is kinda inefficient
        return var{ right + left };
    }
// };

int TestVarVsInt() {
    int varOperationsAmount = 10000000;   // one hundred thousand ops

    auto t1 = chrono::high_resolution_clock::now();
    var vz = 0;
    var vy = 8;
    var vx = 7;
    for (int i = 0; i < varOperationsAmount; i++) {
        vz += vx * vy;
    }
    cout << "vz " << vz << endl;

    int varDuration = chrono::duration_cast<chrono::milliseconds>(std::chrono::high_resolution_clock::now()-t1).count();

    cout << "var type operations took " << varDuration << " milliseconds\n\n";

    int intOperationsAmount = 10000000; // ten million ops

    auto t3 = chrono::high_resolution_clock::now();
    int iz = 0;
    for (int i = 0; i < intOperationsAmount; i++) {
        iz += 7 * 8;
    }
    cout << "iz " << iz << endl;

    int intDuration = chrono::duration_cast<chrono::milliseconds>(std::chrono::high_resolution_clock::now()-t3).count();
    cout << "int type operations took " << intDuration << " milliseconds\n\n";


    // var statsObj = map<string, var>{
    //     { "varOperations", map<string, var>{ 
    //         { "amount", varOperationsAmount },
    //         { "duration", varDuration }},
    //     },
    //     { "intOperations",  map<string, var>{ 
    //         { "amount", varOperationsAmount },
    //         { "duration", intDuration }},
    //     }
    // };

    // cout << "stats: " << statsObj << endl << endl; 

    return 0;
}


// int main() {
//     // var me = map<string, var>{
//     //     { "hey", "its me" },
//     //     { "maybe", false },
//     //     { "floater", 7.88865 },
//     //     { "interino", 65 },
//     //     { "anotherObject", map<string, var>{
//     //         { "intBro", 789 },
//     //         { "fadedString", "woah" }
//     //     }},
//     // };
//     // cout << "var: " << me << endl;
//     var test = "6";
//     test = 6;

//     // test["me"] = {
//     //     "something", 7,
//     //     "this", {
//     //         "you", true,
//     //     }
//     // };
//     // TODO: just use this as a hack right now
//     // test = {};
//     test = { "me", 6 };
//     test["me"] = "7";
    
//     // test = {};
//     cout << "object_test: " << test << " " << endl;
//     // test.deallocate();
//     // // test = "me";
//     // // // test.~var();
//     // printf("test 7\n");
//     // var test = 7;
//     // cout << "test: " << test << " " << endl;
//     // test = 9;
//     // printf("test 8\n");
//     // // test = "yo";
//     // // test = "hi";
//     // cout << "test: " << test << " " << endl;
//     // // test = map<string, var>{
//     // //     { "hey", "its me" },
//     // //     { "maybe", false },
//     // //     { "floater", 7.88865 },
//     // //     { "interino", 65 },
//     // //     { "anotherObject", map<string, var>{
//     // //         { "intBro", 789 },
//     // //         { "fadedString", "woah" }
//     // //     }},
//     // // };
//     // test = "hey";
//     // cout << "tester0: " << test << " " << endl;
//     // test = 7 + 3;
//     // cout << "tester: " << test << " " << endl;
//     // test = .999;
//     // cout << "tester: " << test << " " << endl;
//     // test = false;
//     // cout << "tester: " << test << " " << endl;


//     // // TODO: this needs to create a new object
//     // test["floater"] = 7;
//     // cout << "tester: " << test << " " << endl;
//     // // cout << "test: " << test["floater"] << endl;

//     // // int someVariable = 5;

//     // // TestVarVsInt();

//     return 0;
// }



