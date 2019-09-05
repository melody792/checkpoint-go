<template>
    <div id="customers" class="container">
        <!--弹窗组件 -->   <!-- :message  父组件给子组件传值 -->
        <alert v-if="content" :message="content"></alert>

        <h1 class="page-header">用户管理系统</h1>
        <!--v-model 数据双向绑定 -->
        <input type="text" class="form-control" placeholder="搜索" v-model="search">
        <table class="table table-striped">
            <thead>
            <tr>
                <th>姓名</th>
                <th>电话</th>
                <th>邮箱</th>
                <th></th>
            </tr>
            </thead>
            <tbody>
            <!--循环遍历数组 绑定信息 -->
            <tr v-for="(customer,index) in filterBy(customers,search)" :key="index">
                <td>{{customer.name}}</td>
                <td>{{customer.phone}}</td>
                <td>{{customer.email}}</td>
                <td>

                    <router-link class="btn btn-default" :to="'/detail/'+customer.id">详情</router-link>
                </td>
            </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
    import detail from "./PersonalPage.vue";
    export default {
        components: {
            alert: alert,
            detail: detail
        },
        data() {
            //这里存放数据
            return {
                msg: "我是customers组件",
                customers: [],
                content: "",
                search: ""
            };
        },
        methods: {
            filterBy(customers, value) {
                return customers.filter(function(customer) {
                    return customer.name.match(value);
                    //有就返回，没有就不显示    什么都不输就显示所有
                    //filter（）过滤 只能用于数组
                });
            },
            //获取用户信息
            fetchCustomers() {
                this.$http.get("http://localhost:1001/users").then(res => {
                    this.customers = res.body;
                });
            }
        },

        // 页面加载以后直接调用函数   挂载数据，绑定事件等等
        created() {
            if (this.$route.query.content) {
                this.content = this.$route.query.content;
            }
            this.fetchCustomers();
        },
//信息更新以后，再重新获取用户信息进行绑定
        updated() {
            this.fetchCustomers();
        },

        //生命周期 - 挂载完成（可以访问DOM元素）
        mounted() {}
    };
</script>